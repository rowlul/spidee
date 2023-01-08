package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/urfave/cli/v2"
)

func NewGetCommand() *cli.Command {
	cmd := &cli.Command{
		Name:         internal.CommandGet,
		Usage:        "Get message by specified id",
		ArgsUsage:    "<message id>",
		Action:       actionGet,
		OnUsageError: usageError,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: internal.FlagJSON, Usage: "return JSON message object"},
		},
		HideHelpCommand: true,
	}

	return cmd
}

func actionGet(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)

	arg, err := cmdcontext.Uint64Arg(ctx)
	if err != nil {
		return err
	}

	id := discord.MessageID(arg)
	message, err := client.Message(discord.MessageID(id))
	if err != nil {
		return err
	}

	if ctx.Bool(internal.FlagJSON) {
		json, err := json.Marshal(message)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
