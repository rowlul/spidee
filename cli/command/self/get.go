package self

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/cli/util"
	"github.com/urfave/cli/v2"
)

var GetCommand = cli.Command{
	Name:    "get",
	Usage:   "Get webhook",
	Aliases: []string{"g"},
	Action: func(c *cli.Context) error {
		client := *webhook.New(discord.WebhookID(c.Int("id")), c.String("token"))

		webhook, err := client.Get()
		if err != nil {
			return err
		}

		if c.Bool("json") {
			if !c.Bool("no-redact") {
				webhook.Token = ""
			}

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
		&cli.BoolFlag{Name: "json", Usage: "output webhook object in json"},
		&cli.BoolFlag{Name: "format", Usage: "format output"},
		&cli.BoolFlag{Name: "no-redact", Usage: "don't redact sensitive data, e.g. webhook token"},
	},
}
