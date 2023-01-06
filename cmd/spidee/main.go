package main

import (
	"fmt"
	"os"

	"github.com/rowlul/spidee/internal/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
