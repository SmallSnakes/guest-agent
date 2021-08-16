package backups

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

//获取网络信息
func getNetInfo() ([]interface{}, error) {
	netInterfaces, err := net.Interfaces()
	log.Println(netInterfaces)
	if err != nil {
		return nil, err
	}

	var netInfos []interface{}
	for _, netInterface := range netInterfaces {
		macName := netInterface.Name
		macAddr := netInterface.HardwareAddr.String()
		addrs, err := netInterface.Addrs()
		if err != nil {
			return nil, err
		}

		if len(macAddr) == 0 {
			continue
		}

		var macIP string
		for _, address := range addrs {
			ipNet, isValidIpNet := address.(*net.IPNet)
			if isValidIpNet && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					macIP = ipNet.IP.String()
				}
			}
		}

		netInfo := map[string]string{}
		netInfo["name"] = macName
		netInfo["mac"] = macAddr
		netInfo["ip"] = macIP
		netInfos = append(netInfos, netInfo)
	}

	return netInfos, nil
}

func GetLLDPInfo() (map[string]string, error) {
	cmd := exec.Command("powershell", "/c", "Invoke-DiscoveryProtocolCapture -Type LLDP | Get-DiscoveryProtocolData")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	//strData := strings.Replace(string(data), " ", "", -1)
	strData := strings.TrimSpace(string(data))
	lineDatas := strings.Split(strData, "\r\n")

	LLDPInfos := map[string]string{}
	for _, lineData := range lineDatas {

		if lineData[:1] != " " {
			rowDate := strings.Split(lineData, ":")
			LLDPInfos[strings.TrimSpace(rowDate[0])] = strings.TrimSpace(rowDate[1])
		}

	}

	return LLDPInfos, err
}
