package self

import (
	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/urfave/cli/v2"
)

var DeleteCommand = cli.Command{
	Name:    "delete",
	Usage:   "Delete webhook",
	Aliases: []string{"d"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))
		err := client.Delete()
		return err
	},
}
