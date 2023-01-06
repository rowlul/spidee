package self

import (
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/urfave/cli/v2"
)

func NewDeleteCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   pkg.CommandDelete,
		Usage:  "Delete webhook",
		Action: actionDelete,
	}

	return cmd
}

func actionDelete(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)
	return client.Delete()
}
