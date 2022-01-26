package util

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/urfave/cli/v2"
)

func TestBuildFilesFromContext(t *testing.T) {
	is := is.New(t)

	args := []string{"spidee", "--file", "files_test.go"}
	app := &cli.App{
		Flags: []cli.Flag{&cli.StringSliceFlag{Name: "file"}},
		Action: func(c *cli.Context) error {
			files, err := BuildFilesFromContext(c)
			is.NoErr(err)

			a, err := io.ReadAll(files[0].Reader)
			is.NoErr(err)

			reader, err := os.OpenFile("files_test.go", os.O_RDONLY, os.ModeAppend)
			is.NoErr(err)

			b, err := io.ReadAll(reader)
			is.NoErr(err)

			is.Equal(files[0].Name, filepath.Base(reader.Name()))
			is.Equal(a, b)

			return nil
		},
	}

	app.Run(args)
}
