package command

import (
	"strconv"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/utils/json/option"
	"github.com/urfave/cli/v2"
)

func EditCommand(client webhook.Client) cli.Command {
	return cli.Command{
		Name:    "edit",
		Usage:   "edit message",
		Aliases: []string{"e"},
		Action: func(c *cli.Context) error {
			messageId, err := strconv.Atoi(c.Args().First())
			if err != nil {
				//log.Fatalln(`Required argument "messageId" not set or integer`)
				return err
			}

			return client.EditMessage(
				discord.MessageID(messageId),
				webhook.EditMessageData{Content: option.NewNullableString(c.String("content"))},
			)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "content",
				Usage:   "plain text",
				Aliases: []string{"c"},
			},
		},
	}
}
