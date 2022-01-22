package cli

import (
	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/rowlul/spidee/cli/command"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "spidee",
	Usage: "A command line interface for Discord webhooks",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "id",
			Usage:   "webhook id",
			EnvVars: []string{"SPIDEE_WEBHOOK_ID"},
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "webhook token",
			EnvVars: []string{"SPIDEE_WEBHOOK_TOKEN"},
		},
	},
	Before: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))
		_, err := client.Get()

		var (
			sendCommand   = command.SendCommand(client)
			editCommand   = command.EditCommand(client)
			deleteCommand = command.DeleteCommand(client)
		)

		c.App.Commands = []*cli.Command{
			&sendCommand,
			&editCommand,
			&deleteCommand,
		}

		return err
	},
}