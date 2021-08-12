// +build windows

package utils

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"net"
)

const wqlNetworkAdapter = "SELECT Index,Name, MACAddress, NetworkAddresses,  Installed, PhysicalAdapter  " +
	"FROM Win32_NetworkAdapter"

type win32NetworkAdapter struct {
	Index            int
	Name             string
	MACAddress       string
	NetworkAddresses string
	Installed        bool
	PhysicalAdapter  bool
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
		ipaddr := getIP(net_interface.Index)
		print(ipaddr)
		var netInfo = &NetArg{
			Name:        net_interface.Name,
			MacAddress:  net_interface.MACAddress,
			Ipv4Address: "",
			Ipv6Address: "",
			HasCarrier:  net_interface.Installed,
			//LLDP:        []",
		}
		netInfos = append(netInfos, netInfo)
	}
	return netInfos, err
}

func load() ([]win32NetworkAdapter, error) {
	// Getting info from WMI
	var win32NetDescriptions []win32NetworkAdapter
	if err := wmi.Query(wqlNetworkAdapter, &win32NetDescriptions); err != nil {
		return nil, err
	}
	return win32NetDescriptions, nil
}

func getIP(inx int) []net.Addr{
	byIndex, err := net.InterfaceByIndex(inx)
	if err != nil {
		return nil
	}
	addresses, _ := byIndex.Addrs()
	fmt.Println(addresses)

	return addresses
}
