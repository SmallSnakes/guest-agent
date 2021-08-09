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
	memInfo["tatal"] = info.Total
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var memory uint64
	proc.Call(uintptr(unsafe.Pointer(&memory)))
	memInfo["physical_mb"] = memory / 1024

	return memInfo, err
}

// func GetMemoryWinInfo(c *gin.Context) {
// 	var info interface{}
// 	var err interface{}

// 	info, err = GetMemoryWin()
// 	if err != nil {
// 		log.Println("get cpu info error ", err)
// 	}

// 	c.JSON(200, gin.H{
// 		"memory": info,
// 	})
// }
