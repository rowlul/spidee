package self

import (
	"encoding/json"
	"fmt"

	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/urfave/cli/v2"
)

func NewGetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   pkg.CommandGet,
		Usage:  "Get webhook",
		Action: actionGet,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: pkg.FlagJSON, Usage: "return JSON message object"},
		},
	}

	return cmd
}

func actionGet(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	webhook, err := client.Get()
	if err != nil {
		return err
	}

	if ctx.Bool(pkg.FlagJSON) {
		json, err := json.Marshal(webhook)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
