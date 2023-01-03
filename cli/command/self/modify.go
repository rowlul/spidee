package self

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/rowlul/spidee/cli/util"
	"github.com/urfave/cli/v2"
)

var ModifyCommand = cli.Command{
	Name:    "modify",
	Usage:   "Modify webhook",
	Aliases: []string{"m"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		if len(c.String("payload")) > 0 {
			payload := c.String("payload")
			data := api.ModifyWebhookData{}
			json.Unmarshal([]byte(payload), &data)

			webhook, err := client.Modify(data)

			if c.Bool("json") {
				format := c.Bool("format")

				msg, err := util.StringifyObject(webhook, format)
				if err != nil {
					return err
				}

				fmt.Println(msg)
				return nil
			}

			return err
		}

		if len(c.String("username")) == 0 && len(c.String("avatar")) != 0 {
			return errors.New("username must be supplied along with avatar")
		} else if len(c.String("username")) == 0 {
			return errors.New("no username or avatar supplied")
		}

		data := api.ModifyWebhookData{
			Name: option.NewString(c.String("username")),
			// channel id cannot be set due to webhook client limitations
			// https://github.com/discordjs/discord.js/issues/7518
		}

		if len(c.String("avatar")) > 0 {
			avatar, err := util.BuildImageFromContext(c)
			if err != nil {
				return err
			}

			data.Avatar = avatar
		}

		webhook, err := client.Modify(data)
		if err != nil {
			return err
		}

		if c.Bool("json") {
			format := c.Bool("format")

			msg, err := util.StringifyObject(webhook, format)
			if err != nil {
				return err
			}

			fmt.Println(msg)
			return nil
		}

		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "username",
			Usage:   "webhook username",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:      "avatar",
			Usage:     "path to webhook avatar image",
			Aliases:   []string{"a"},
			TakesFile: true,
		},
		&cli.StringFlag{
			Name:    "payload",
			Usage:   "raw json payload",
			Aliases: []string{"p"}},
		&cli.BoolFlag{Name: "json", Usage: "output webhook object in json"},
		&cli.BoolFlag{Name: "format", Usage: "format output"},
	},
}
