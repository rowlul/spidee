package main

import (
	"log"
	"os"

	"github.com/rowlul/spidee/cli"
)

func main() {
	log.SetFlags(0)

	err := cli.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
