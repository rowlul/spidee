package self

import "github.com/urfave/cli/v2"

var ModifyCommand = cli.Command{
	Name:    "modify",
	Usage:   "Modify webhook",
	Aliases: []string{"m"},
}
