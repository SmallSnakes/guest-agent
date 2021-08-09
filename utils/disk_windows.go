package utils

import (
	"github.com/StackExchange/wmi"
	"log"
)

type Win32_DiskDrive struct {
	Caption      string
	Name         string
	DeviceID     string
	Model        string
	Index        int
	Partitions   int
	Size         int
	PNPDeviceID  string
	Status       string
	SerialNumber string
	Manufacturer string
	MediaType    string
	Description  string
	SystemName   string
}

func GetDiskDrivers() []Win32_DiskDrive {

	var dst []Win32_DiskDrive

	query := wmi.CreateQuery(&dst, "")
	if err := wmi.Query(query, &dst); err != nil {
		log.Println(err.Error())
	}

	//for key, value := range dst {
	//  log.Println(fmt.Sprintf(`Disk%d: %v`, key+1, value))
	//}

	return dst
}
