package main

import (
	"log"
	"os"
	"strconv"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/utils/json/option"
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
				},
				Action: func(c *cli.Context) error {
					return client.Execute(webhook.ExecuteData{
						Content:   content,
						Username:  username,
						AvatarURL: avatarUrl,
						TTS:       tts},
					)

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
