package command

import (
	"errors"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/urfave/cli/v2"
)

var DeleteCommand = cli.Command{
	Name:      "delete",
	Usage:     "Delete message",
	ArgsUsage: "<message id>",
	Aliases:   []string{"d"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		messageId, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return errors.New("message id not set or not integer")
		}

		err = client.DeleteMessage(discord.MessageID(messageId))
		if err != nil {
			return err
		}

		return nil
	},
}
