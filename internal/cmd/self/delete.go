package self

import (
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   internal.CommandDelete,
		Usage:  "Delete webhook",
		Action: actionDelete,
	}

	return cmd
}

func actionDelete(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)
	return client.Delete()
}
