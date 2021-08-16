// +build !windows

package backups

import (
	"fmt"
	"github.com/jaypipes/ghw"
)

type NetArg struct {
	Name        string   `json:"name"`
	MacAddress  string   `json:"mac_address"`
	Ipv4Address string   `json:"ipv4_address"`
	Ipv6Address string   `json:"ipv6_address"`
	HasCarrier  bool     `json:"has_carrier"`
	LLDP        []string `json:"lldp"`
}

// GetNetInfo 获取网络信息
func GetNetInfo() ([]interface{}, error) {
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	var netInfos []interface{}
	for _, nic := range net.NICs {
		if nic.IsVirtual{
			continue
		}
		var netInfo = &NetArg{
			Name: nic.Name,
			MacAddress: nic.MacAddress,
			//Ipv4Address: nic.Capabilities,
			Ipv6Address: nic.String(),

		}
		netInfos = append(netInfos,netInfo)
	}
	return netInfos, err
}
