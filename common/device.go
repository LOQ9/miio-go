package common

import "miio-go/subscription"

type DeviceInfo struct {
	Life                int    `json:"life"`
	ConfigTime          int    `json:"cfg_time"`
	BindKey             string `json:"bindkey"`
	UID                 int    `json:"uid"`
	FirmwareVersion     string `json:"fw_ver"`
	HardwareVersion     string `json:"hw_ver"`
	MacAddress          string `json:"mac"`
	Model               string `json:"model"`
	Token               string `json:"token"`
	WifiFirmwareVersion string `json:"wifi_fw_ver"`
}

type Device interface {
	subscription.SubscriptionTarget

	ID() uint32
	GetLabel() (string, error)
	GetInfo() (DeviceInfo, error)
	GetToken() []byte
}
