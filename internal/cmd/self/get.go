package self

import (
	"encoding/json"
	"fmt"

	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/urfave/cli/v2"
)

func NewGetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   internal.CommandGet,
		Usage:  "Get webhook",
		Action: actionGet,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: internal.FlagJSON, Usage: "return JSON message object"},
			&cli.BoolFlag{Name: "no-redact", Usage: "don't redact sensitive data, e.g. webhook token"},
		},
	}

	return cmd
}

func actionGet(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)

	webhook, err := client.Get()
	if err != nil {
		return err
	}

	if ctx.Bool(internal.FlagJSON) {
		if !ctx.Bool(internal.FlagNoRedact) {
			webhook.Token = ""
		}

		json, err := json.Marshal(webhook)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
