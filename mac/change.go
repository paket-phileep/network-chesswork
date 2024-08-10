package mac

import (
	"net"

	"github.com/vishvananda/netlink"
)

func ChangeMAC(name string, macaddr string) error {

	iface, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	hwaddr, err := net.ParseMAC(macaddr)
	if err != nil {
		return err
	}
	err = netlink.LinkSetHardwareAddr(iface, hwaddr)
	if err != nil {
		return err
	}
	
	return nil
}
