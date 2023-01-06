package cmd

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand() *cli.Command {
	cmd := &cli.Command{
		Name:      pkg.CommandDelete,
		Usage:     "Delete webhook message",
		ArgsUsage: "<message id>",
		Action:    actionDelete,
	}

	return cmd
}

func actionDelete(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	arg, err := context.Uint64Arg(ctx)
	if err != nil {
		return err
	}

	id := discord.MessageID(arg)
	return client.DeleteMessage(id)
}
