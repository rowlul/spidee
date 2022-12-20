package command

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/cli/util"
	"github.com/urfave/cli/v2"
)

var SendCommand = cli.Command{
	Name:    "send",
	Usage:   "Send message",
	Aliases: []string{"s"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		if util.IsStdin() {
			s := strings.Join(util.ReadStdin(), "\n")
			data := webhook.ExecuteData{}
			if json.Valid([]byte(s)) {
				json.Unmarshal([]byte(s), &data)
			} else {
				data.Content = s
			}

			if c.Bool("wait") {
				message, err := client.ExecuteAndWait(data)
				if err != nil {
					return err
				}

				if c.Bool("json") {
					format := c.Bool("format")

					msg, err := util.StringifyMessage(message, format)
					if err != nil {
						return nil
					}

					log.Println(msg)
				}
			} else if c.Bool("json") {
				log.Fatalln("error: must wait for message before output (use --wait)")
			}

			err := client.Execute(data)
			return err
		}

		if len(c.String("payload")) > 0 {
			payload := c.String("payload")
			data := webhook.ExecuteData{}
			json.Unmarshal([]byte(payload), &data)

			if c.Bool("wait") {
				message, err := client.ExecuteAndWait(data)
				if err != nil {
					return err
				}

				if c.Bool("json") {
					format := c.Bool("format")

					msg, err := util.StringifyMessage(message, format)
					if err != nil {
						return nil
					}

					log.Println(msg)
				}
			} else if c.Bool("json") {
				log.Fatalln("error: must wait for message before output (use --wait)")
			}

			err := client.Execute(data)
			return err
		}

		if len(c.String("content")) == 0 &&
			len(c.StringSlice("file")) == 0 &&
			!c.Bool("embed") {
			cli.ShowSubcommandHelp(c)
			log.Fatalln("error: no content, file, or embed supplied")
		}

		files, err := util.BuildFilesFromContext(c)
		if err != nil {
			log.Fatalln(util.FormatFileError(err))
		}

		embeds, err := util.BuildEmbedsFromContext(c)
		if err != nil {
			log.Fatalln(util.FormatEmbedError(err))
		}

		data := webhook.ExecuteData{
			Content:   c.String("content"),
			Username:  c.String("username"),
			AvatarURL: c.String("avatar-url"),
			TTS:       c.Bool("tts"),
			Files:     files,
		}

		if c.Bool("embed") {
			data.Embeds = embeds
		}

		if c.Bool("wait") {
			message, err := client.ExecuteAndWait(data)
			if err != nil {
				return err
			}

			if c.Bool("json") {
				format := c.Bool("format")

				msg, err := util.StringifyMessage(message, format)
				if err != nil {
					return nil
				}

				log.Println(msg)
			}
		} else if c.Bool("json") {
			log.Fatalln("error: must wait for message before output (use --wait)")
		}

		err = client.Execute(data)
		return err
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "content",
			Usage:   "plain text",
			Aliases: []string{"c"},
		},
		&cli.StringFlag{
			Name:    "username",
			Usage:   "webhook username",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "avatar-url",
			Usage:   "webhook avatar url",
			Aliases: []string{"a"},
		},
		&cli.BoolFlag{
			Name:    "tts",
			Usage:   "narrate message",
			Aliases: []string{"t"},
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
		&cli.BoolFlag{
			Name:    "wait",
			Usage:   "wait for message to be created",
			Aliases: []string{"w"}},
		&cli.BoolFlag{Name: "json", Usage: "output message object in json"},
		&cli.BoolFlag{Name: "format", Usage: "format output"},
	},
}
