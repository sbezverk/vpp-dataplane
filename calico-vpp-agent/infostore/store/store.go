package store

import (
	"fmt"
	"strings"
	"sync"

	"github.com/projectcalico/vpp-dataplane/calico-vpp-agent/infostore"
	"github.com/sirupsen/logrus"
)

type info struct {
	log *logrus.Entry
	sync.Mutex
	// The key is a combination of pod's namespace and pod's name
	store map[string]*infostore.Record
}

var _ infostore.Manager = &info{}

func (i *info) AddPodInfo(r *infostore.Record) error {
	i.Lock()
	defer i.Unlock()
	if r == nil {
		return fmt.Errorf("pod record is nil")
	}
	key := makeKey(r.Name, r.Namespace)

	if _, ok := i.store[key]; ok {
		i.log.Warnf("pod %s/%s is already in the store", r.Namespace, r.Name)
	}
	i.store[key] = r
	i.log.Infof("pod %s/%s (key: %s) is added to the store, total entries in the store: %d", r.Namespace, r.Name, key, len(i.store))
	return nil
}

func (i *info) RemovePodInfo(interfaceName string) error {
	i.Lock()
	defer i.Unlock()
	found := false
	key := ""
	for _, v := range i.store {
		if strings.Compare(v.InterfaceName, interfaceName) == 0 {
			found = true
			key = makeKey(v.Name, v.Namespace)
			break
		}
	}
	if !found {
		return fmt.Errorf("pod with interface name %s is not found in the store", interfaceName)
	}
	r, ok := i.store[key]
	if !ok {
		i.log.Warnf("pod with the key %s is not found in the store", key)
		return nil
	}
	delete(i.store, key)
	i.log.Infof("pod %s/%s is deleted from the store, total entries in the store: %d", r.Namespace, r.Name, len(i.store))

	return nil
}

// GetPodInfo returns a Record of information for a specific pod
func (i *info) GetPodInfo(n, ns string) (*infostore.Record, error) {
	i.Lock()
	defer i.Unlock()
	key := makeKey(n, ns)
	r, ok := i.store[key]
	if !ok {
		i.log.Errorf("pod with the key %s is not found in the store", key)
		return nil, fmt.Errorf("pod with the key %s is not found in the store", key)
	}

	return r, nil
}

// NewInfoStore creates a new instance of the store
func NewInfoStore(l *logrus.Entry) infostore.Manager {
	return &info{
		store: make(map[string]*infostore.Record),
		log:   l,
	}
}

func makeKey(n, ns string) string {
	return ns + "_" + n
}
