// +build !windows

package utils

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
	var netInfos []interface{}
	return netInfos, err
}

