package command

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/rowlul/spidee/cli/util"
	"github.com/urfave/cli/v2"
)

var EditCommand = cli.Command{
	Name:    "edit",
	Usage:   "Edit message",
	Aliases: []string{"e"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		messageId, err := strconv.Atoi(c.Args().First())
		if err != nil {
			cli.ShowSubcommandHelp(c)
			log.Fatalln("error: message id not set or not integer")
		}

		if util.IsStdin() {
			s := strings.Join(util.ReadStdin(), "\n")
			data := webhook.EditMessageData{}
			if json.Valid([]byte(s)) {
				json.Unmarshal([]byte(s), &data)
			} else {
				data.Content = option.NewNullableString(s)
			}

			_, err := client.EditMessage(discord.MessageID(messageId), data)
			return err
		}

		if len(c.String("payload")) > 0 {
			payload := c.String("payload")
			data := webhook.EditMessageData{}
			json.Unmarshal([]byte(payload), &data)

			_, err := client.EditMessage(discord.MessageID(messageId), data)
			return err
		}

		if len(c.String("content")) == 0 &&
			len(c.StringSlice("file")) == 0 &&
			!c.Bool("embed") && !util.IsStdin() {
			cli.ShowSubcommandHelpAndExit(c, 2)
		}

		files, err := util.BuildFilesFromContext(c)
		if err != nil {
			log.Fatalln(util.FormatFileError(err))
		}

		embeds, err := util.BuildEmbedsFromContext(c)
		if err != nil {
			log.Fatalln(util.FormatEmbedError(err))
		}

		data := webhook.EditMessageData{
			Content: option.NewNullableString(c.String("content")),
			Files:   files,
		}

		if c.Bool("embed") {
			data.Embeds = &embeds
		}

		_, err = client.EditMessage(
			discord.MessageID(messageId),
			data,
		)

		return err
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Usage:   "plain text",
			Aliases: []string{"c"},
		},
		&cli.StringSliceFlag{
			Name:      "file",
			Usage:     "webhook attachment",
			Aliases:   []string{"f"},
			TakesFile: true,
		},
		&cli.BoolFlag{
			Name:    "embed",
			Usage:   "include embed",
			Aliases: []string{"e"},
		},
		&cli.StringFlag{Name: "embed-title", Usage: "embed title"},
		&cli.StringFlag{Name: "embed-description", Usage: "embed description"},
		&cli.StringFlag{Name: "embed-url", Usage: "embed url"},
		&cli.StringFlag{Name: "embed-timestamp", Usage: "embed timestamp (now|RFC3339 timestamp)"},
		&cli.IntFlag{Name: "embed-color", Usage: "embed color (hex)"},
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
		&cli.StringSliceFlag{Name: "embed-field", Usage: "embed field (name,value,inline)"},
		&cli.StringFlag{
			Name:    "payload",
			Usage:   "raw json payload",
			Aliases: []string{"p"}},
	},
}
