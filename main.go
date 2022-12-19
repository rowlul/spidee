package main

import (
	"os"

	"github.com/rowlul/spidee/cli"
	"github.com/rowlul/spidee/cli/util"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		util.HandleError(err)
	}
}
