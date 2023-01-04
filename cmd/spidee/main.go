package main

import (
	"fmt"
	"os"

	"github.com/rowlul/spidee/pkg/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
