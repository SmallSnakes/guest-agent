// +build !windows

package utils

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

func GetMemory() (interface{}, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	memInfo := map[string]interface{}{}
	memInfo["tatal"] = info.Total

	cmd := exec.Command("sh", "-c", "dmidecode --type 17 | grep Size")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	var physical_mb int64
	strData := strings.TrimSpace(string(data))
	lineDatas := strings.Split(strData, "\n")
	for _, lineData := range lineDatas {
		if find := strings.Contains(lineData, "GB"); find {
			rowDate := strings.Split(lineData, ":")
			memArr := strings.Split(strings.TrimSpace(rowDate[1]), " ")
			mem, _ := strconv.ParseInt(memArr[0], 10, 64)
			physical_mb = physical_mb + (mem * 1024)
		}
		memInfo["physical_mb"] = physical_mb
	}

	return memInfo, err
}
