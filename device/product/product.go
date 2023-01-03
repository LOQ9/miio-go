package product

import "fmt"

type Product uint16

const (
	PowerPlug Product = iota << 1
	Yeelight
	AirPurifierV7
	ZhimiAirpurifierMA2
	Unknown
)

func GetModel(modelName string) (Product, error) {
	switch modelName {
	case "chuangmi.plug.m1":
		return PowerPlug, nil
	case "yeelink.light.color1":
		return Yeelight, nil
	case "zhimi.airpurifier.v7":
		return AirPurifierV7, nil
	case "zhimi.airpurifier.ma2":
		return ZhimiAirpurifierMA2, nil
	default:
		return Unknown, fmt.Errorf("Unknown product for device type %s", modelName)
	}
}
