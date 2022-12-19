package main

import (
	"log"
	"os"

	"github.com/rowlul/spidee/cli"
	"github.com/rowlul/spidee/cli/util"
)

func main() {
	log.SetFlags(0)

	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatalln(util.FormatError(err))
	}
}
