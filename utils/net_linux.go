// +build !windows

package utils

import "fmt"

type win32NetworkAdapter struct {
	GUID            string
	Name            string
	Installed       bool
	PhysicalAdapter bool
	ProductName     string
	Manufacturer    string
}

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
fmt.Println()
	var netInfo NetArg
	for _, device := range getAllDevice() {
		switchInfo, localInfo := LLDPInfo(device.Name)
		netInfo.Name = localInfo.Name
		netInfo.MacAddress = localInfo.MacAddress
		netInfo.Ipv4Address = localInfo.Ipv4Address
		netInfo.Ipv6Address = localInfo.Ipv6Address
		netInfo.LLDP = *switchInfo
		netInfos = append(netInfos, netInfo)
	}
	return netInfos, nil
}

