package commands

import (
	"miio-go/model"

	"github.com/alecthomas/kingpin"
)

const (
	OutFormatDefault = "default"
	OutFormatJSON    = "json"
)

type Command struct {
	App          *kingpin.Application
	SharedClient *model.Client
	OutFormat    string
}

func NewCommand(app *kingpin.Application, sharedClient *model.Client, outFormat string) *Command {
	return &Command{
		App:          app,
		SharedClient: sharedClient,
		OutFormat:    outFormat,
	}
}

func (c *Command) SetSharedClient(sharedClient *model.Client) {
	c.SharedClient = sharedClient
}
