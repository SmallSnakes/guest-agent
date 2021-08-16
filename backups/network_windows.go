// +build windows

package backups

import (
	"fmt"
	"net"
	"strings"
)

const wqlNetworkAdapter = "SELECT InterfaceIndex ,Name, MACAddress, NetworkAddresses, Installed, PhysicalAdapter,  " +
	"ProductName FROM Win32_NetworkAdapter"

type win32NetworkAdapter struct {
	InterfaceIndex   int
	Name             string
	MACAddress       string
	NetworkAddresses string
	Installed        bool
	PhysicalAdapter  bool
	ProductName		 string
}

type NetArg struct {
	Name        string `json:"name"`
	MacAddress  string `json:"mac_address"`
	Ipv4Address string `json:"ipv4_address"`
	Ipv6Address string `json:"ipv6_address"`
	HasCarrier  bool   `json:"has_carrier"`
	//LLDP        []string `json:"lldp"`
}

func GetNetInfo() ([]interface{}, error) {
	nets, err := load()
	if err != nil {
		return nil, err
	}

	var netInfos []interface{}
	for _, net_interface := range nets {
		if !net_interface.PhysicalAdapter {
			continue
		}

		// 获取ipv4和ipv6
		var ipv4 string
		var ipv6 string
		ipaddr := getIP(net_interface.InterfaceIndex)
		for _, addr := range ipaddr {
			if isIPv4(addr.String()) {
				ipv4 = addr.String()
			} else if isIPv6(addr.String()) {
				ipv6 = addr.String()
			}
		}
		fmt.Println(net_interface.ProductName)
		var netInfo = &NetArg{
			Name:        net_interface.Name,
			MacAddress:  net_interface.MACAddress,
			Ipv4Address: ipv4,
			Ipv6Address: ipv6,
			HasCarrier:  net_interface.Installed,
			//LLDP:        []",
		}
		netInfos = append(netInfos, netInfo)
	}
	return netInfos, err
}

// Getting info from WMI
func load() ([]win32NetworkAdapter, error) {
	var win32NetDescriptions []win32NetworkAdapter
	if err := wmi.Query(wqlNetworkAdapter, &win32NetDescriptions); err != nil {
		return nil, err
	}
	return win32NetDescriptions, nil
}

//获取ipaddr
func getIP(inx int) []net.Addr {
	byIndex, err := net.InterfaceByIndex(inx)
	if err != nil {
		return nil
	}
	addresses, _ := byIndex.Addrs()

	return addresses
}

//判断是否为ipv4
func isIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

//判断是否为ipv6
func isIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}
