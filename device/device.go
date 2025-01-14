package device

import (
	"time"

	"miio-go/common"
	"miio-go/device/product"

	"miio-go/protocol/packet"
	"miio-go/protocol/transport"
)

type Device interface {
	common.Device

	Handle(*packet.Packet) error
	Close() error
	Seen() time.Time
	Provisional() bool
	SetProvisional(bool)
	GetProduct() (product.Product, error)
	Discover() error
	RefreshThrottle() <-chan struct{}
	Outbound() transport.Outbound
}
