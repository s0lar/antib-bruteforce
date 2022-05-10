package netlist

import (
	"net"
	"sync"
)

type NetList struct {
	mx   sync.RWMutex
	list map[string]*net.IPNet
}

func NewNetList() *NetList {
	return &NetList{
		list: make(map[string]*net.IPNet),
	}
}

//	Load - for fast load lists of CIDR
func (n *NetList) Load(cidrList []string) error {
	for _, cidr := range cidrList {
		err := n.Add(cidr)
		if err != nil {
			return err
		}
	}
	return nil
}

//	Add - save CIDR in list
func (n *NetList) Add(cidr string) error {
	_, IPNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}

	n.mx.Lock()
	defer n.mx.Unlock()
	//	IPNet already in list
	if _, ok := n.list[cidr]; ok {
		return nil
	}

	n.list[cidr] = IPNet

	return nil
}

//	Find - trying to find IP in list
func (n *NetList) Find(ip string) bool {
	checkIP := net.ParseIP(ip)

	for _, IPNet := range n.list {
		if IPNet.Contains(checkIP) {
			return true
		}
	}
	return false
}

//	Find - trying to find IP in list
func (n *NetList) Remove(cidr string) error {
	if _, _, err := net.ParseCIDR(cidr); err != nil {
		return err
	}

	n.mx.Lock()
	defer n.mx.Unlock()

	delete(n.list, cidr)
	return nil
}
