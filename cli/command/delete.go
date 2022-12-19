package command

import (
	"strconv"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/urfave/cli/v2"
)

var DeleteCommand = cli.Command{
	Name:    "delete",
	Usage:   "delete a message",
	Aliases: []string{"d"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		messageId, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return err
		}

		return client.DeleteMessage(discord.MessageID(messageId))
	},
}
