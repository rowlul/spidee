package self

import (
	"github.com/rowlul/spidee/internal"
	"github.com/urfave/cli/v2"
)

func NewSelfCommand() *cli.Command {
	cmd := &cli.Command{
		Name:         internal.CommandSelf,
		Usage:        "Refer to webhook",
		OnUsageError: usageError,
		Subcommands: []*cli.Command{
			NewGetCommand(),
			NewModifyCommand(),
			NewDeleteCommand(),
		},
	}

	return cmd
}

func usageError(ctx *cli.Context, err error, isSubcommand bool) error {
	cli.ShowSubcommandHelp(ctx)
	return err
}
