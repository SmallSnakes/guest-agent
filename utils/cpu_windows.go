// +build windows

package utils

import (
	"github.com/StackExchange/wmi"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
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
	const CPUAdapter = "SELECT AddressWidth,Architecture FROM Win32_Processor"
	type win32CPUAdapterr struct {
		AddressWidth uint16
		Architecture int16
	}

	var win32CPUDescriptions []win32CPUAdapterr
	err := wmi.Query(CPUAdapter, &win32CPUDescriptions)
	if err != nil {
		return "", err
	}
	var architectures = map[int16]string{
		0: "x86",
		1: "MIPS",
		2: "Alpha",
		3: "PowerPC",
		5: "ARM",
		6: "ia64",
		9: "x64",
	}
	addressWidth := strconv.Itoa(int(win32CPUDescriptions[0].AddressWidth))
	architecture := architectures[win32CPUDescriptions[0].Architecture]
	return architecture + "_" + addressWidth, nil
}
