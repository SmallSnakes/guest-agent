// +build windows

package utils

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"strings"
)

type win32NetworkAdapter struct {
	GUID              string
	Name              string
	Installed         bool
	PhysicalAdapter   bool
	ProductName       string
	Manufacturer string}

type LLDPArg struct {
	ChassisId       string `json:"chassis_id"`
	PortId          string `json:"port_id"`
	TTL             uint16 `json:"ttl"`
	PortDescription string `json:"port_description"`
	SYSName         string `json:"sys_name"`
	SYSDescription  string `json:"sys_description"`
	SYSCapabilities bool   `json:"sys_capabilities"`
	MGMTAddress     string `json:"mgmt_address"`
	ORGSpecific     string `json:"org_specific"`
}
type NetArg struct {
	Name        string  `json:"name"`
	MacAddress  string  `json:"mac_address"`
	Ipv4Address string  `json:"ipv4_address"`
	Ipv6Address string  `json:"ipv6_address"`
	HasCarrier  bool    `json:"has_carrier"`
	LLDP        LLDPArg `json:"lldp"`
	Vendor      string  `json:"vendor"`
	Product     string  `json:"product"`
	ClientId    string  `json:"client_id"`
	Biosdevname string  `json:"biosdevname"`
}

func GetNetInfo() ([]interface{}, error) {
	var netInfos []interface{}

	phyNic, err := IsPhysicalNet()
	if err != nil {
		return nil, err
	}

	for _, adapter := range phyNic {
		var netInfo NetArg
		for _, device := range getAllDevice() {
			if strings.Contains(device.Name, adapter.GUID) {
				switchInfo, localInfo := LLDPInfo(device.Name)
				netInfo.Name = localInfo.Name
				netInfo.MacAddress = localInfo.MacAddress
				netInfo.Ipv4Address = localInfo.Ipv4Address
				netInfo.Ipv6Address = localInfo.Ipv6Address
				netInfo.HasCarrier = adapter.Installed
				netInfo.LLDP = *switchInfo
				netInfo.Vendor = adapter.Manufacturer
				netInfo.Product = adapter.ProductName

			}
		}
		netInfos = append(netInfos, netInfo)
	}

	return netInfos, nil
}

func IsPhysicalNet() ([]win32NetworkAdapter, error) {
	const wqlNetworkAdapter = "SELECT GUID,Name, ProductName,Installed,Manufacturer, PhysicalAdapter FROM Win32_NetworkAdapter"
	var win32NetDescriptions []win32NetworkAdapter
	if err := wmi.Query(wqlNetworkAdapter, &win32NetDescriptions); err != nil {
		return nil, err
	}

	var PhysicalNet []win32NetworkAdapter
	for _, description := range win32NetDescriptions {
		if description.PhysicalAdapter == false {
			continue
		}
		fmt.Println(description)
		PhysicalNet = append(PhysicalNet, description)
	}
	return PhysicalNet, nil
}
