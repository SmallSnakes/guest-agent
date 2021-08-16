package utils

import (
	"fmt"
	"github.com/jaypipes/ghw"
)

type LLDPArg struct {
	ChassisId       string `json:"chassis_id"`
	PortId          string `json:"port_id"`
	TTL             string `json:"ttl"`
	PortDescription string `json:"port_description"`
	SYSName         string `json:"sys_name"`
	SYSDescription  string `json:"sys_description"`
	SYSCapabilities string `json:"sys_capabilities"`
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
}

func GetNetInfo() ([]interface{}, error) {
	var netInfos []interface{}

	//for _, device := range getAllDevice() {
	//	//不严谨，暂时没想到更好的过滤办法
	//	if len(device.Addresses) == 0 || strings.Contains(
	//		device.Description, "Virtual") {
	//		continue
	//	}
	//	for _, address := range device.Addresses {
	//		fmt.Println(address.IP)
	//	}
	//
	//	//netInfo := &NetArg{
	//	//	Name: device.Name,
	//	//	MacAddress: device.Addresses[],
	//	//
	//	//}
	//}
	//LLDPInfo(device.Name)
	return netInfos,nil
}


func isVirtual() {
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}
	for _, net := range net.NICs {
		if net.IsVirtual {
			continue
		}
		fmt.Println(net)
	}
}
