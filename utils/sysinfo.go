package utils

import "github.com/shirou/gopsutil/host"

func GetHostInfo() string {
	info, _ := host.Info()
	return info.OS
}
