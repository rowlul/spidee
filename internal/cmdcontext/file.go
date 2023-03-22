package cmdcontext

import (
	"os"
	"path/filepath"

	"github.com/diamondburned/arikawa/v3/utils/sendpart"
	"github.com/rowlul/spidee/internal"
	"github.com/urfave/cli/v2"
)

// Files reads files from string slice flag and returns a slice of sendpart.File
// with each file named respectively.
func Files(ctx *cli.Context) ([]sendpart.File, error) {
	var files []sendpart.File

	paths := ctx.StringSlice(internal.FlagFile)
	for _, f := range paths {
		file, err := os.OpenFile(f, os.O_RDONLY, os.ModeAppend)
		if err != nil {
			return files, err
		}
		files = append(files, sendpart.File{Name: filepath.Base(file.Name()), Reader: file})
	}

	return files, nil
}
