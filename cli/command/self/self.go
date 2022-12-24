package self

import "github.com/urfave/cli/v2"

var SelfCommand = cli.Command{
	Name:        "self",
	Usage:       "Refer to webhook",
	Subcommands: []*cli.Command{},
}
