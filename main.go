package main

import (
	"os"

	"github.com/rowlul/spidee/cli"
	"github.com/rowlul/spidee/cli/util"
)

func main() {
	err := cli.App.Run(os.Args)
	if err != nil {
		util.HandleError(err)
	}
}
