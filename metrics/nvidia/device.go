package nvidia

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type Device struct {
	Index              uint8  `json:"index"`
	Name               string `json:"name"`
	MemoryInfoTotal    uint64 `json:"memory-info-total"`
	MemoryInfoFree     uint64 `json:"memory-info-free"`
	MemoryInfoUsed     uint64 `json:"memory-info-used"`
	MemoryUtilization  uint8  `json:"memory-utilization"`
	DecoderUtilization uint8  `json:"decoder-utilization"`
	EncoderUtilization uint8  `json:"encoder-utilization"`
	GPUUtilization     uint8  `json:"gpu-utilization"`

	Temp          uint32 `json:"temp"`
	PowerUsage    uint32 `json:"power-usage"`
	ClockGraphics uint32 `json:"clock-graphics"`
	ClockSm       uint32 `json:"clock-sm"`
	ClockMem      uint32 `json:"clock-mem"`
	FanSpeed      uint32 `json:"fan-speed"`

	PCIInfoBusID          string `json:"pci-info-bus-id"`            //!< The tuple domain:bus:device.function PCI identifier (&amp; NULL terminator)
	PCIInfoDomain         uint32 `json:"pci-info-domain"`            //!< The PCI domain on which the device's bus resides, 0 to 0xffff
	PCIInfoBus            uint32 `json:"pci-info-bus"`               //!< The bus on which the device resides, 0 to 0xff
	PCIInfoDevice         uint32 `json:"pci-info-device"`            //!< The device's id on the bus, 0 to 31
	PCIInfoPCIDeviceID    uint32 `json:"pci-info-pci-device-id"`     //!< The combined 16-bit device id and 16-bit vendor id
	PCIInfoPCISubSystemID uint32 `json:"pci-info-pci-sub-system-id"` //!< The 32-bit Sub System Device ID

	Processes Processes `json:"processes"`
}

func (it *Device) String() string {
	return fmt.Sprintf("{"+
		"id/index: %d, "+
		"name: %s, "+
		"bud-id: %s, "+
		"mem(free/used/total): %s/%s/%s, "+
		"mem(pfree/pused): %.1f%%/%.1f%%, "+
		"enc/dec/gpu/mem: %d%%/%d%%/%d%%/%d%%, "+
		"processes: %d"+
		"}",
		it.Index,
		it.Name,
		it.PCIInfoBusID,
		humanize.Bytes(it.MemoryInfoFree), humanize.Bytes(it.MemoryInfoUsed), humanize.Bytes(it.MemoryInfoTotal),
		it.MemoryInfoPFree(), it.MemoryInfoPUsed(),
		it.EncoderUtilization, it.DecoderUtilization, it.GPUUtilization, it.MemoryUtilization,
		len(it.Processes),
	)
}

func (it *Device) MemoryInfoPFree() float64 {
	if it.MemoryInfoTotal == 0 {
		return 0
	}

	return (float64(it.MemoryInfoFree) / float64(it.MemoryInfoTotal)) * 100
}

func (it *Device) MemoryInfoPUsed() float64 {
	if it.MemoryInfoTotal == 0 {
		return 0
	}

	return (float64(it.MemoryInfoUsed) / float64(it.MemoryInfoTotal)) * 100
}
