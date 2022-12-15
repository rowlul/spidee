package command

import (
	"strconv"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/urfave/cli/v2"
)

func DeleteCommand(client webhook.Client) cli.Command {
	return cli.Command{
		Name:    "delete",
		Usage:   "delete a message",
		Aliases: []string{"d"},
		Action: func(c *cli.Context) error {
			messageId, err := strconv.Atoi(c.Args().First())
			if err != nil {
				return err
			}

			return client.DeleteMessage(discord.MessageID(messageId))
		},
	}
}
