package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/utils/json/option"
	"github.com/diamondburned/arikawa/v2/utils/sendpart"
	"github.com/urfave/cli/v2"
)

func main() {
	var webhookId int
	var webhookToken string

	var content, username, avatarUrl string
	var tts bool

	var client webhook.Client

	app := &cli.App{
		Name:  "spidee",
		Usage: "A command line interface for Discord webhooks",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "id",
				Usage:       "webhook id",
				EnvVars:     []string{"SPIDEE_WEBHOOK_ID"},
				Destination: &webhookId,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "webhook token",
				EnvVars:     []string{"SPIDEE_WEBHOOK_TOKEN"},
				Destination: &webhookToken,
			},
		},
		Before: func(c *cli.Context) error {
			// initialize client and test webhook before proceeding
			client = *webhook.New(discord.WebhookID(webhookId), webhookToken)
			_, err := client.Get()

			return err
		},
		Commands: []*cli.Command{
			{
				Name:    "send",
				Usage:   "send a message",
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "content",
						Usage:       "plain text",
						Aliases:     []string{"c"},
						Destination: &content,
					},
					&cli.StringFlag{
						Name:        "username",
						Usage:       "webhook username",
						Aliases:     []string{"u"},
						Destination: &username,
					},
					&cli.StringFlag{
						Name:        "avatar-url",
						Usage:       "webhook avatar url",
						Aliases:     []string{"a"},
						Destination: &avatarUrl,
					},
					&cli.BoolFlag{
						Name:        "tts",
						Usage:       "narrate message",
						Aliases:     []string{"t"},
						Destination: &tts,
					},
					&cli.StringSliceFlag{
						Name:    "file",
						Usage:   "webhook attachment",
						Aliases: []string{"f"},
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
				Action: func(c *cli.Context) error {
					var files []sendpart.File
					for _, f := range c.StringSlice("file") {
						file, err := os.OpenFile(f, os.O_RDONLY, os.ModeAppend)
						if err != nil {
							return err
						}

						files = append(files, sendpart.File{Name: file.Name(), Reader: file})
					}

					embedTimestamp, err := time.Parse(time.RFC3339, c.String("embed-timestamp"))
					if c.String("embed-timestamp") == "now" {
						embedTimestamp = time.Now()
					} else if len(c.String("embed-timestamp")) > 0 && err != nil {
						log.Fatalln(`Argument "timestamp" must be valid RFC3339 timestamp`)
					}

					var embedFields []discord.EmbedField
					for _, f := range c.StringSlice("embed-field") {
						field := strings.Split(f, ",")

						var inline bool
						if len(field) > 2 {
							inline, err = strconv.ParseBool(field[2])
							if err != nil {
								log.Fatalln(`Argument "inline" must be bool`)
							}
						}

						embedFields = append(embedFields, discord.EmbedField{Name: field[0], Value: field[1], Inline: inline})
					}

					embeds := []discord.Embed{{
						Title:       c.String("embed-title"),
						Description: c.String("embed-description"),
						URL:         c.String("embed-url"),
						Timestamp:   discord.NewTimestamp(embedTimestamp),
						Color:       discord.Color(c.Int("embed-color")),
						Footer:      &discord.EmbedFooter{Text: c.String("embed-footer-text"), Icon: c.String("embed-footer-icon")},
						Image:       &discord.EmbedImage{URL: c.String("embed-image-url")},
						Thumbnail:   &discord.EmbedThumbnail{URL: c.String("embed-thumbnail-url")},
						Video:       &discord.EmbedVideo{URL: c.String("embed-video-url")},
						Provider:    &discord.EmbedProvider{Name: c.String("embed-provider-name"), URL: c.String("embed-provider-url")},
						Author:      &discord.EmbedAuthor{Name: c.String("embed-author-name"), URL: c.String("embed-author-url"), Icon: c.String("embed-author-icon")},
						Fields:      embedFields},
					}

					err = embeds[0].Validate()
					if err != nil {
						log.Fatalln(err)
					}

					data := webhook.ExecuteData{
						Content:   content,
						Username:  username,
						AvatarURL: avatarUrl,
						TTS:       tts,
						Files:     files,
						Embeds:    embeds,
					}

					err = client.Execute(data)
					if err != nil && strings.Contains(err.Error(), "Invalid Form Body") { // unpleasant workaround to send message if no embed supplied
						data.Embeds = nil
						err = client.Execute(data)
					}

					return err
				},
			},
			{
				Name:    "edit",
				Usage:   "edit a message",
				Aliases: []string{"e"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "content",
						Usage:       "plain text",
						Aliases:     []string{"c"},
						Destination: &content,
					},
				},
				Action: func(c *cli.Context) error {
					messageId, err := strconv.Atoi(c.Args().First())
					if err != nil {
						log.Fatalln(`Required argument "messageId" not set or integer`)
					}

					return client.EditMessage(
						discord.MessageID(messageId),
						webhook.EditMessageData{Content: option.NewNullableString(content)},
					)
				},
			},
			{
				Name:    "delete",
				Usage:   "delete a message",
				Aliases: []string{"d"},
				Action: func(c *cli.Context) error {
					messageId, err := strconv.Atoi(c.Args().First())
					if err != nil {
						log.Fatalln(`Required argument "messageId" not set or integer`)
					}

					return client.DeleteMessage(discord.MessageID(messageId))
				},
			},
		},
	}

	log.SetFlags(0)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
