package self

import "github.com/urfave/cli/v2"

var GetCommand = cli.Command{
	Name:    "get",
	Usage:   "Get webhook",
	Aliases: []string{"g"},
}
