package cmd

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand() *cli.Command {
	cmd := &cli.Command{
		Name:      internal.CommandDelete,
		Usage:     "Delete webhook message",
		ArgsUsage: "<message id>",
		Action:    actionDelete,
	}

	return cmd
}

func actionDelete(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)

	arg, err := cmdcontext.Uint64Arg(ctx)
	if err != nil {
		return err
	}

	id := discord.MessageID(arg)
	return client.DeleteMessage(id)
}
