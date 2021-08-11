package utils

import (
	"github.com/jaypipes/ghw"
)

func GetMemory() (interface{}, error) {
	memory, err := ghw.Memory()
	if err != nil {
		return nil,err
	}
	memInfo := map[string]interface{}{}
	memInfo["total"] = memory.TotalUsableBytes
	memInfo["physical_mb"] = memory.TotalPhysicalBytes/1024/1024
	return memInfo,err
}