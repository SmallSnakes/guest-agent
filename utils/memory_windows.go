// +build windows

package utils

import (
	"syscall"
	"unsafe"

	"github.com/shirou/gopsutil/mem"
)

func GetMemory() (interface{}, error) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	memInfo := map[string]interface{}{}
	memInfo["total"] = info.Total
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var memory uint64
	_, _, err = proc.Call(uintptr(unsafe.Pointer(&memory)))
	if err != nil {
		return nil, err
	}
	memInfo["physical_mb"] = memory / 1024

	return memInfo, err
}
