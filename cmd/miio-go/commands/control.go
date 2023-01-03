package commands

import (
	"fmt"
	"time"

	"miio-go/capability"
	"miio-go/common"
	"miio-go/device"

	"github.com/alecthomas/kingpin"
)

var sharedDevice common.Device

func (c *Command) findDevice(timeout time.Duration) (common.Device, error) {
	timeoutCh := time.After(timeout)
	sub, err := c.SharedClient.NewSubscription()
	if err != nil {
		panic(err)
	}
	events := sub.Events()
	defer sub.Close()

	for {
		select {
		case event := <-events:
			switch event.(type) {
			case common.EventNewDevice:
				dev := event.(common.EventNewDevice).Device

				return dev, nil

				// if dev.ID() == deviceId {
				// 	return dev, nil
				// }
			}
		case <-timeoutCh:
			return nil, fmt.Errorf("timed out attempting to connect to available devices")
		}
	}
}

func (c *Command) InstallControl() {
	controlCmd := c.App.Command("control", "Control devices")
	// deviceId := controlCmd.Flag("device-id", "The ID of the device to control").Required().Uint32()

	controlCmd.Action(func(ctx *kingpin.ParseContext) (err error) {
		sharedDevice, err = c.findDevice(time.Second * 5)
		return
	})

	c.installAirPurifierV7(controlCmd)
}

func (c *Command) installAirPurifierV7(parent *kingpin.CmdClause) {
	cmd := parent.Command("airpurifier", "Set Air Purifier")

	cmd.Action(func(ctx *kingpin.ParseContext) error {
		var airPurifier *capability.AirPurifierV7

		switch sharedDevice.(type) {
		case *device.AirPurifierV7:
			airPurifier = sharedDevice.(*device.AirPurifierV7).AirPurifierV7
		default:
			return fmt.Errorf("device with type %T cannot have Air Purifier", sharedDevice)
		}

		return airPurifier.Info()
	})
}
