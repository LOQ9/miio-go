package device

import (
	"miio-go/capability"
)

type AirPurifierV7 struct {
	Device
	*capability.AirPurifierV7
}

func NewAirPurifierV7(device Device) *AirPurifierV7 {
	return &AirPurifierV7{
		Device:        device,
		AirPurifierV7: capability.NewAirPurifierV7(device, device.Outbound()),
	}
}
