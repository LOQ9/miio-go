package device

import (
	"fmt"

	"miio-go/device/product"
)

// Classify determines the underlying product of the device and returns an
// appropriate device implementation.
func Classify(dev Device) (Device, error) {
	if !dev.Provisional() {
		return dev, nil
	}

	p, err := dev.GetProduct()
	if err != nil {
		return nil, err
	}

	defer dev.SetProvisional(false)

	switch p {
	case product.AirPurifierV7, product.ZhimiAirpurifierMA2:
		return NewAirPurifierV7(dev), nil
	default:
		return nil, fmt.Errorf("Classify: Unknown device type")
	}
}
