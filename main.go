package main

import (
	"os"

	"github.com/rowlul/spidee/cli"
	"github.com/rowlul/spidee/cli/util"
)

func main() {
	var err error

	if util.IsStdin() {
		err = cli.App.Run([]string{os.Args[0], "send"})
	} else {
		err = cli.App.Run(os.Args)
	}

	if err != nil {
		util.HandleError(err)
	}
}
