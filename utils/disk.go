package utils

import (
	"github.com/jaypipes/ghw"
)

type DiskArg struct {
	Name               string `json:"name" :"name"`
	Model              string `json:"model" :"model"`
	SizeBytes          uint64 `json:"size" :"size_bytes"`
	Rotational         bool   `json:"rotational" :"rotational"`
	WWN                string `json:"wwn" :"wwn"`
	SerialNumber       string `json:"serial" :"serial_number"`
	Vendor             string `json:"vendor" :"vendor"`
	WwnWithExtension   string `json:"wwn_with_extension" :"wwn_with_extension"`
	WwnVendorExtension string `json:"wwn_vendor_extension" :"wwn_vendor_extension"`
	Hctl               string `json:"hctl" :"hctl"`
	BusPath            string `json:"by_path"`
}

func GetDiskInfo() (interface{}, error) {
	block, err := ghw.Block()
	if err != nil {
		return nil, err
	}

	var diskInfos []interface{}
	for _, disk := range block.Disks {
		if disk.Model == "unknown" {
			continue
		}

		var isRotational bool
		if string(disk.DriveType.String()) == "SSD"{
			isRotational = false
		}else if string(disk.DriveType.String()) == "HDD" {
			isRotational = true
		}

		diskInfo := &DiskArg{
			Name:               disk.Name,
			Model:              disk.Model,
			SizeBytes:          disk.SizeBytes,
			Rotational:         isRotational,
			WWN:                disk.WWN,
			SerialNumber:       disk.SerialNumber,
			Vendor:             disk.Vendor,
			WwnWithExtension:   disk.WWN,
			WwnVendorExtension: "",
			Hctl:               "",
			BusPath:            disk.BusPath,
		}
		diskInfos = append(diskInfos, diskInfo)
	}
	return diskInfos, nil

}
