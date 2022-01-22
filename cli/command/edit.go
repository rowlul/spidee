package command

import (
	"strconv"
	"strings"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/utils/json/option"
	"github.com/rowlul/spidee/cli/util"
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

			embeds, err := util.BuildEmbedsFromContext(c)
			if err != nil {
				return err
			}

			data := webhook.EditMessageData{
				Content: option.NewNullableString(c.String("content")),
				Embeds:  &embeds,
			}

			err = client.EditMessage(
				discord.MessageID(messageId),
				data,
			)
			if err != nil && strings.Contains(err.Error(), "Invalid Form Body") { // unpleasant workaround to send message if no embed supplied
				data.Embeds = nil
				err = client.EditMessage(
					discord.MessageID(messageId),
					data,
				)
			}

			return err
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "content",
				Usage:   "plain text",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{Name: "embed-title", Usage: "embed title"},
			&cli.StringFlag{Name: "embed-description", Usage: "embed description"},
			&cli.StringFlag{Name: "embed-url", Usage: "embed url"},
			&cli.StringFlag{Name: "embed-timestamp", Usage: "embed timestamp"},
			&cli.IntFlag{Name: "embed-color", Usage: "embed color"},
			&cli.StringFlag{Name: "embed-footer-text", Usage: "embed footer text"},
			&cli.StringFlag{Name: "embed-footer-icon", Usage: "embed footer icon"},
			&cli.StringFlag{Name: "embed-image-url", Usage: "embed image url"},
			&cli.StringFlag{Name: "embed-thumbnail-url", Usage: "embed thumbnail url"},
			&cli.StringFlag{Name: "embed-video-url", Usage: "embed video url"},
			&cli.StringFlag{Name: "embed-provider-name", Usage: "embed provider name"},
			&cli.StringFlag{Name: "embed-provider-url", Usage: "embed provider url"},
			&cli.StringFlag{Name: "embed-author-name", Usage: "embed author name"},
			&cli.StringFlag{Name: "embed-author-url", Usage: "embed author url"},
			&cli.StringFlag{Name: "embed-author-icon", Usage: "embed author icon"},
			&cli.StringSliceFlag{Name: "embed-field", Usage: "embed field"},
		},
	}
}
