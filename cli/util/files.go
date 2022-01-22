package util

import (
	"os"

	"github.com/diamondburned/arikawa/v2/utils/sendpart"
	"github.com/urfave/cli/v2"
)

func BuildFilesFromContext(c *cli.Context) ([]sendpart.File, error) {
	var err error

	var files []sendpart.File
	for _, f := range c.StringSlice("file") {
		var file *os.File
		file, err = os.OpenFile(f, os.O_RDONLY, os.ModeAppend)
		files = append(files, sendpart.File{Name: file.Name(), Reader: file})
	}

	return files, err
}
