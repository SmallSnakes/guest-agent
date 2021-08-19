package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type LocalNetArg struct {
	Name        string `json:"name"`
	MacAddress  string `json:"mac_address"`
	Ipv4Address string `json:"ipv4_address"`
	Ipv6Address string `json:"ipv6_address"`
}

func getAllDevice() []pcap.Interface {
	var devices []pcap.Interface
	devices, _ = pcap.FindAllDevs()
	return devices
}

func LLDPInfo(deviceName string) (LLDPArg, LocalNetArg) {
	handle, err := pcap.OpenLive(deviceName, int32(65535), true, -1*time.Second)
	fmt.Println(handle.ReadPacketData())
	if err != nil {
		log.Println(err)
	}
	defer handle.Close()

	switchInfo := LLDPArg{}
	localInfo := LocalNetArg{}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	portIDSubtypes := map[int]string{1: "Interface Name", 2: "Local"}
	packet := packetSource.Packets()
	for packet := range packet {
		if lldpPacket, ok := packet.Layer(layers.LayerTypeLinkLayerDiscovery).(*layers.LinkLayerDiscovery); ok {
			lldpInfo, _ := packet.Layer(layers.LayerTypeLinkLayerDiscoveryInfo).(*layers.LinkLayerDiscoveryInfo)
			if lldpPacket.PortID.Subtype.String() == "Interface Name" {
				switchInfo = interfaceSwitch(*lldpPacket, *lldpInfo)
				delete(portIDSubtypes, 1)
			} else if lldpPacket.PortID.Subtype.String() == "Local" {
				localInfo = local(*lldpPacket)
				delete(portIDSubtypes, 2)
			}

			if len(portIDSubtypes) == 0 {
				break
			}
		}
	}
	return switchInfo, localInfo

}

//解析switch信息
func interfaceSwitch(lldpPacket layers.LinkLayerDiscovery, lldpInfo layers.LinkLayerDiscoveryInfo) LLDPArg {
	info := LLDPArg{}
	info.ChassisId = BytesTo16(lldpPacket.ChassisID.ID)
	info.PortId = BytesToString(lldpPacket.PortID.ID)
	info.TTL = lldpPacket.TTL
	info.PortDescription = lldpInfo.PortDescription
	info.SYSName = lldpInfo.SysName
	info.SYSDescription = lldpInfo.SysDescription
	info.SYSCapabilities = lldpInfo.SysCapabilities.SystemCap.Other
	info.MGMTAddress = (net.IP)(lldpInfo.MgmtAddress.Address).String()
	info.ORGSpecific = BytesTo16(lldpInfo.OrgTLVs[0].Info)
	return info
}

//解析本地信息
func local(lldpPacket layers.LinkLayerDiscovery) LocalNetArg {
	localInfo := LocalNetArg{}
	var ipv4 string
	var ipv6 string
	for _, value := range lldpPacket.Values {
		if value.Type.String() == "Management Address" {
			if value.Length == 12 {
				ipv4 = (net.IP)(value.Value[2 : len(value.Value)-6]).String()
			} else if value.Length == 24 {
				ipv6 = (net.IP)(value.Value[2 : len(value.Value)-6]).String()
			}

		}
	}
	localInfo.Name = BytesToString(lldpPacket.PortID.ID)
	localInfo.MacAddress = ParseNetNameMac(BytesTo16(lldpPacket.ChassisID.ID))
	localInfo.Ipv4Address = ipv4
	localInfo.Ipv6Address = ipv6
	return localInfo
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

// 转换标准mac
func ParseNetNameMac(name string) string {
	buf, err := hex.DecodeString(name)
	if err != nil {
		log.Println(err)
	}
	return net.HardwareAddr(buf).String()
}
