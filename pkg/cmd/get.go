package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/urfave/cli/v2"
)

func NewGetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:      pkg.CommandGet,
		Usage:     "Get message by specified id",
		ArgsUsage: "<message id>",
		Action:    actionGet,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: pkg.FlagJSON, Usage: "return JSON message object"},
		},
	}

	return cmd
}

func actionGet(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	arg, err := context.Uint64Arg(ctx)
	if err != nil {
		return err
	}

	id := discord.MessageID(arg)
	message, err := client.Message(discord.MessageID(id))
	if err != nil {
		return err
	}

	if ctx.Bool(pkg.FlagJSON) {
		json, err := json.Marshal(message)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
