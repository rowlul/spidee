package cli

import (
	"github.com/rowlul/spidee/cli/command"
	"github.com/urfave/cli/v2"
)

var Version string

func NewApp() *cli.App {
	return &cli.App{
		Name:    "spidee",
		Usage:   "A command line interface for Discord webhooks",
		Version: Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "id",
				Usage:    "webhook id",
				EnvVars:  []string{"SPIDEE_WEBHOOK_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "token",
				Usage:    "webhook token",
				EnvVars:  []string{"SPIDEE_WEBHOOK_TOKEN"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			&command.SendCommand,
			&command.EditCommand,
			&command.DeleteCommand,
		},
		DisableSliceFlagSeparator: true,
		DefaultCommand:            "send",
	}
}
