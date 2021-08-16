package utils

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"reflect"
	"strings"
	"time"
	"unsafe"
)


func getAllDevice() []pcap.Interface {
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	return devices
}

func LLDPInfo() {

	handle, err := pcap.OpenLive("\\Device\\NPF_{80AFA1A5-2B1C-4E31-9B85-6A4557002FD2}", int32(65535), true, -1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	portIDSubtypes := map[int]string{1: "Interface Name", 2: "Local"}
	packet := packetSource.Packets()
	for packet := range packet {
		if lldpPacket, ok := packet.Layer(layers.LayerTypeLinkLayerDiscovery).(*layers.LinkLayerDiscovery); ok {
			lldpInfo,_ := packet.Layer(layers.LayerTypeLinkLayerDiscoveryInfo).(*layers.LinkLayerDiscoveryInfo)
			if lldpPacket.PortID.Subtype.String() == "Interface Name" {
				fmt.Println("chassis",BytesTo16(lldpPacket.ChassisID.ID ))
				fmt.Println("ttl",lldpPacket.TTL )
				fmt.Println("PortDescription",lldpInfo.PortDescription )
				fmt.Println("SysName",lldpInfo.SysName )
				fmt.Println("SysDescription",lldpInfo.SysDescription )
				fmt.Println("SysCapabilities",lldpInfo.SysCapabilities.SystemCap )
				fmt.Println("MgmtAddress",lldpInfo.MgmtAddress.Address )
				fmt.Println("OrgTLVs",lldpInfo.OrgTLVs )

				delete(portIDSubtypes, 1)
			} else if lldpPacket.PortID.Subtype.String() == "Local" {
				//fmt.Println("macaddres",BytesTo16(got.ChassisID.ID ))
				delete(portIDSubtypes, 2)
			}
			if len(portIDSubtypes) == 0 {
				break
			}
		}

	}

}

// BytesToString []byte 转 string
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

// BytesTo16 []byte 转 16进制
func BytesTo16(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}
