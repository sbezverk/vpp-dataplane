// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.09-rc0~361-gab9444728

// Package calico contains generated bindings for API file calico.api.
//
// Contents:
//   2 messages
//
package calico

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/projectcalico/vpp-dataplane/vpplink/binapi/20.09-rc0~361-gab9444728/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "calico"
	APIVersion = ""
	VersionCrc = 0xa6f24e44
)

// CalicoEnableDisableInterfaceSnat defines message 'calico_enable_disable_interface_snat'.
type CalicoEnableDisableInterfaceSnat struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIP6     bool                           `binapi:"bool,name=is_ip6" json:"is_ip6,omitempty"`
	IsEnable  bool                           `binapi:"bool,name=is_enable" json:"is_enable,omitempty"`
}

func (m *CalicoEnableDisableInterfaceSnat) Reset() { *m = CalicoEnableDisableInterfaceSnat{} }
func (*CalicoEnableDisableInterfaceSnat) GetMessageName() string {
	return "calico_enable_disable_interface_snat"
}
func (*CalicoEnableDisableInterfaceSnat) GetCrcString() string { return "7e99d008" }
func (*CalicoEnableDisableInterfaceSnat) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CalicoEnableDisableInterfaceSnat) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIP6
	size += 1 // m.IsEnable
	return size
}
func (m *CalicoEnableDisableInterfaceSnat) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsIP6)
	buf.EncodeBool(m.IsEnable)
	return buf.Bytes(), nil
}
func (m *CalicoEnableDisableInterfaceSnat) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIP6 = buf.DecodeBool()
	m.IsEnable = buf.DecodeBool()
	return nil
}

// CalicoEnableDisableInterfaceSnatReply defines message 'calico_enable_disable_interface_snat_reply'.
type CalicoEnableDisableInterfaceSnatReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CalicoEnableDisableInterfaceSnatReply) Reset() { *m = CalicoEnableDisableInterfaceSnatReply{} }
func (*CalicoEnableDisableInterfaceSnatReply) GetMessageName() string {
	return "calico_enable_disable_interface_snat_reply"
}
func (*CalicoEnableDisableInterfaceSnatReply) GetCrcString() string { return "e8d4e804" }
func (*CalicoEnableDisableInterfaceSnatReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CalicoEnableDisableInterfaceSnatReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CalicoEnableDisableInterfaceSnatReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CalicoEnableDisableInterfaceSnatReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_calico_binapi_init() }
func file_calico_binapi_init() {
	api.RegisterMessage((*CalicoEnableDisableInterfaceSnat)(nil), "calico_enable_disable_interface_snat_7e99d008")
	api.RegisterMessage((*CalicoEnableDisableInterfaceSnatReply)(nil), "calico_enable_disable_interface_snat_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CalicoEnableDisableInterfaceSnat)(nil),
		(*CalicoEnableDisableInterfaceSnatReply)(nil),
	}
}