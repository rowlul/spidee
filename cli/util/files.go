package util

import (
	"os"

	"github.com/diamondburned/arikawa/v2/utils/sendpart"
	"github.com/urfave/cli/v2"
)

func BuildFilesFromContext(c *cli.Context) ([]sendpart.File, error) {
	var files []sendpart.File
	for _, f := range c.StringSlice("file") {
		file, err := os.OpenFile(f, os.O_RDONLY, os.ModeAppend)
		if err != nil {
			return files, err
		}
		files = append(files, sendpart.File{Name: file.Name(), Reader: file})
	}

	return files, nil
}
