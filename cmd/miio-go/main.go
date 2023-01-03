package main

import (
	"os"

	"miio-go/cmd/miio-go/commands"
	"miio-go/common"
	"miio-go/model"

	"github.com/alecthomas/kingpin"

	"github.com/sirupsen/logrus"
)

func main() {
	app := kingpin.New("miio-go CLI", "CLI application to manually test miio-go functionality")
	local := app.Flag("local", "Send broadcast to 127.0.0.1 instead of 255.255.255.255 (For use with locally hosted simulator)").Bool()
	address := app.Flag("address", "Set MiiO address").Required().String()
	token := app.Flag("token", "Set MiiO token").Required().String()
	deviceId := app.Flag("device-id", "Set DeviceID").Int()
	logLevel := app.Flag("log-level", "Set MiiO to a specific log level").Default("warn").Enum("debug", "warn", "info", "error")
	outFormat := app.Flag("out-format", "Set MiiO output format").Default("default").Enum(commands.OutFormatDefault, commands.OutFormatJSON)

	command := commands.NewCommand(app, nil, *outFormat)

	command.InstallControl()
	command.InstallDiscovery()

	app.Action(func(ctx *kingpin.ParseContext) error {
		level, _ := logrus.ParseLevel(*logLevel)
		l := logrus.New()
		l.SetLevel(level)
		common.SetLogger(l)

		sharedClient, err := model.CreateClient(*local, *address, *token, *deviceId)
		command.SetSharedClient(sharedClient)

		return err
	})

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
