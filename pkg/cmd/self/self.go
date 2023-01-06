package self

import (
	"github.com/rowlul/spidee/pkg"
	"github.com/urfave/cli/v2"
)

func NewSelfCommand() *cli.Command {
	cmd := &cli.Command{
		Name:  pkg.CommandSelf,
		Usage: "Refer to webhook",
		Subcommands: []*cli.Command{
			NewGetCommand(),
			NewModifyCommand(),
			NewDeleteCommand(),
		},
	}

	return cmd
}
