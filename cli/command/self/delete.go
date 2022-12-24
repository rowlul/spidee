package self

import "github.com/urfave/cli/v2"

var DeleteCommand = cli.Command{
	Name:    "delete",
	Usage:   "Delete webhook",
	Aliases: []string{"d"},
}
