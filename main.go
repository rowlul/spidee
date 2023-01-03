package main

import (
	"log"
	"os"

	"github.com/rowlul/spidee/cli"
	"github.com/rowlul/spidee/cli/util"
)

func main() {
	app := cli.NewApp()

	if util.IsStdin() {
		app.DefaultCommand = "send"
	}

	err := app.Run(os.Args)
	if err != nil {
		log.SetFlags(0)
		log.SetPrefix("error: ")
		log.Fatal(err)
	}
}
