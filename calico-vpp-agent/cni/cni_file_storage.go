// Copyright (C) 2020 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cni

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/lunixbochs/struc"
	"github.com/pkg/errors"
	pb "github.com/projectcalico/vpp-dataplane/calico-vpp-agent/cni/proto"
	"github.com/projectcalico/vpp-dataplane/calico-vpp-agent/config"
)

type LocalIPNet struct {
	MaskSize int    `struc:"int8,sizeof=Mask"`
	IP       net.IP `struc:"[16]byte"`
	Mask     net.IPMask
}

type LocalIP struct {
	IP net.IP `struc:"[16]byte"`
}

func (n *LocalIPNet) String() string {
	ipnet := net.IPNet{
		IP:   n.IP,
		Mask: n.Mask,
	}
	return ipnet.String()
}

func (n *LocalIP) String() string {
	return n.IP.String()
}

func (n *LocalIPNet) UpdateSizes() {
	n.MaskSize = len(n.Mask)
}

func (ps *LocalPodSpec) UpdateSizes() {
	ps.RoutesSize = len(ps.Routes)
	ps.ContainerIpsSize = len(ps.ContainerIps)
	ps.InterfaceNameSize = len(ps.InterfaceName)
	ps.NetnsNameSize = len(ps.NetnsName)
	for _, n := range ps.Routes {
		n.UpdateSizes()
	}
}

func (ps *LocalPodSpec) Key() string {
	return fmt.Sprintf("%s--%s", ps.NetnsName, ps.InterfaceName)
}

func (ps *LocalPodSpec) String() string {
	lst := ps.ContainerIps
	strLst := make([]string, 0, len(lst))
	for _, e := range lst {
		strLst = append(strLst, e.String())
	}
	return fmt.Sprintf("%s: %s", ps.Key(), strings.Join(strLst, ", "))
}

type LocalPodSpec struct {
	InterfaceNameSize int `struc:"int16,sizeof=InterfaceName"`
	InterfaceName     string
	NetnsNameSize     int `struc:"int16,sizeof=NetnsName"`
	NetnsName         string
	AllowIpForwarding bool
	RoutesSize        int `struc:"int16,sizeof=Routes"`
	Routes            []LocalIPNet
	ContainerIpsSize  int `struc:"int16,sizeof=ContainerIps"`
	ContainerIps      []LocalIP
	// Pod identifiers
	OrchestratorID string
	WorkloadID     string
	EndpointID     string
}

func (ps *LocalPodSpec) GetRoutes() (routes []*net.IPNet) {
	routes = make([]*net.IPNet, 0, len(ps.Routes))
	for _, r := range ps.Routes {
		routes = append(routes, &net.IPNet{
			IP:   r.IP,
			Mask: r.Mask,
		})
	}
	return routes
}

func (ps *LocalPodSpec) GetContainerIps() (containerIps []*net.IPNet) {
	containerIps = make([]*net.IPNet, 0, len(ps.ContainerIps))
	for _, containerIp := range ps.ContainerIps {
		containerIps = append(containerIps, &net.IPNet{
			IP:   containerIp.IP,
			Mask: getMaxCIDRMask(containerIp.IP),
		})
	}
	return containerIps
}

func (ps *LocalPodSpec) Hasv46() (hasv4 bool, hasv6 bool) {
	hasv4 = false
	hasv6 = false
	for _, containerIP := range ps.ContainerIps {
		if containerIP.IP.To4() == nil {
			hasv6 = true
		} else {
			hasv4 = true
		}
	}
	return hasv4, hasv6
}

func (ps *LocalPodSpec) NeedsSnat(s *Server) (needsSnat bool) {
	needsSnat = false
	for _, containerIP := range ps.GetContainerIps() {
		needsSnat = needsSnat || s.IPNetNeedsSNAT(containerIP)
	}
	return needsSnat
}

func NewLocalPodSpecFromAdd(request *pb.AddRequest) (*LocalPodSpec, error) {
	podSpec := LocalPodSpec{
		InterfaceName:     request.GetInterfaceName(),
		NetnsName:         request.GetNetns(),
		AllowIpForwarding: request.GetSettings().GetAllowIpForwarding(),
		Routes:            make([]LocalIPNet, 0),
		ContainerIps:      make([]LocalIP, 0),

		OrchestratorID: request.Workload.Orchestrator,
		WorkloadID:     request.Workload.Name,
		EndpointID:     request.Workload.Endpoint,
	}
	for _, routeStr := range request.GetContainerRoutes() {
		_, route, err := net.ParseCIDR(routeStr)
		if err != nil {
			return nil, errors.Wrapf(err, "Cannot parse container route %s", routeStr)
		}
		podSpec.Routes = append(podSpec.Routes, LocalIPNet{
			IP:   route.IP,
			Mask: route.Mask,
		})
	}
	for _, requestContainerIP := range request.GetContainerIps() {
		containerIp, _, err := net.ParseCIDR(requestContainerIP.GetAddress())
		if err != nil {
			return nil, fmt.Errorf("Cannot parse address: %s", requestContainerIP.GetAddress())
		}
		// We ignore the prefix len set on the address,
		// for a tun it doesn't make sense
		podSpec.ContainerIps = append(podSpec.ContainerIps, LocalIP{IP: containerIp})
	}

	return &podSpec, nil
}

func NewLocalPodSpecFromDel(request *pb.DelRequest) *LocalPodSpec {
	return &LocalPodSpec{
		InterfaceName: request.GetInterfaceName(),
		NetnsName:     request.GetNetns(),
	}
}

func (s *Server) persistCniServerState() (err error) {
	var buf bytes.Buffer
	for _, podSpec := range s.podInterfaceMap {
		err := struc.Pack(&buf, podSpec)
		if err != nil {
			return errors.Wrap(err, "Error packing data")
		}
	}
	err = ioutil.WriteFile(config.CniServerStateTempFile, buf.Bytes(), 0200)
	if err != nil {
		return errors.Wrapf(err, "Error writing file %s", config.CniServerStateTempFile)
	}
	err = os.Rename(config.CniServerStateTempFile, config.CniServerStateFile)
	if err != nil {
		return errors.Wrapf(err, "Error moving file %s", config.CniServerStateTempFile)
	}
	return nil
}

func (s *Server) loadCniServerState() (podSpecs []*LocalPodSpec, err error) {
	podSpecs = make([]*LocalPodSpec, 0)
	data, err := ioutil.ReadFile(config.CniServerStateFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading file %s", config.CniServerStateFile)
	}
	buf := bytes.NewBuffer(data)
	i := 0
	for buf.Len() > 0 {
		podSpec := &LocalPodSpec{}
		err = struc.Unpack(buf, podSpec)
		if err != nil {
			return nil, errors.Wrapf(err, "Error unpacking")
		}
		podSpecs = append(podSpecs, podSpec)
		i++
		if i > 20 {
			return podSpecs, nil
		}
	}
	return podSpecs, nil
}
