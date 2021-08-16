// +build !windows

package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"os/exec"
	"strconv"
	"strings"
)

type CPUArg struct {
	ModelName    string   `json:"model_name"`
	Frequency    string   `json:"frequency"`
	Count        int      `json:"count"`
	Architecture string   `json:"architecture"`
	Flags        []string `json:"flags"`
}

func GetCPUInfo() (interface{}, error) {

	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	architecture, err := cpuArchitecture()
	if err != nil {
		return nil, err
	}
	cpuInfo := &CPUArg{
		ModelName:    info[0].ModelName,
		Frequency:    strconv.FormatFloat(info[0].Mhz, 'f', -1, 64),
		Count:        len(info),
		Architecture: architecture,
		Flags:        info[0].Flags,
	}
	return cpuInfo, err
}

func cpuArchitecture() (string, error) {
	cmd := exec.Command("bash","-c","lscpu | grep Architecture")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "nil", err
	}
	strData := strings.TrimSpace(string(data))
	rowDate := strings.Split(strData, ":")
	return strings.TrimSpace(rowDate[1]), nil
}
