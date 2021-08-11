package utils

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
	"strconv"
)

func GetCPUInfo() (interface{}, error) {

	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuInfo := map[string]interface{}{}
	cpuInfo["model_name"] = info[0].ModelName
	cpuInfo["frequency"] = strconv.FormatFloat(info[0].Mhz, 'f', -1, 64)
	cpuInfo["count"] = len(info)

	// lscpu 所获取的跟这个不一样
	cpuInfo["architecture"] = runtime.GOARCH
	cpuInfo["flags"] = info[0].Flags
	return cpuInfo, err
}

func Test()  {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	fmt.Println("%v\n", cpu)
	for _, proc := range cpu.Processors {
		fmt.Printf(" %v\n", proc.Cores)
	}

}