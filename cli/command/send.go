package command

import (
	"strings"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/rowlul/spidee/cli/util"
	"github.com/urfave/cli/v2"
)

func SendCommand(client webhook.Client) cli.Command {
	return cli.Command{
		Name:    "send",
		Usage:   "send message",
		Aliases: []string{"s"},
		Action: func(c *cli.Context) error {
			if len(c.String("content")) == 0 && len(c.StringSlice("file")) == 0 && !c.Bool("embed") && !util.IsStdin() {
				cli.ShowCommandHelpAndExit(c, "send", 2)
			}

			files, err := util.BuildFilesFromContext(c)
			if err != nil {
				return err
			}

			embeds, err := util.BuildEmbedsFromContext(c)
			if err != nil {
				return err
			}

			data := webhook.ExecuteData{
				Content:   c.String("content"),
				Username:  c.String("username"),
				AvatarURL: c.String("avatar-url"),
				TTS:       c.Bool("tts"),
				Files:     files,
			}

			if util.IsStdin() {
				data.Content = strings.Join(util.ReadStdin(), "\n")
			}

			if c.Bool("embed") {
				data.Embeds = embeds
			}

			return client.Execute(data)
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
				Name:    "file",
				Usage:   "webhook attachment",
				Aliases: []string{"f"},
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
		},
	}
}
