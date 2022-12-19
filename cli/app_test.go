package cli

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/rowlul/spidee/cli/command"
	"github.com/urfave/cli/v2"
)

func getFlags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    "id",
			Usage:   "webhook id",
			EnvVars: []string{"SPIDEE_WEBHOOK_ID"},
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "webhook token",
			EnvVars: []string{"SPIDEE_WEBHOOK_TOKEN"},
		},
	}
}

func TestSendCommandWithContent(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	args := os.Args[0:1]
	args = append(args, "send", "--content", "content")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithFile(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	args := os.Args[0:1]
	args = append(args, "send", "--file", "app_test.go")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbed(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	url := "https://go.dev/"

	args := os.Args[0:1]
	args = append(args, "send", "--embed",
		"--embed-title", "title", "--embed-description", "description", "--embed-url", url, "--embed-color", "990000",
		"--embed-footer-text", "text", "--embed-footer-icon", url, "--embed-image-url", url, "--embed-thumbnail-url", url,
		"--embed-video-url", url, "--embed-provider-name", "name", "--embed-provider-url", url, "--embed-author-name", "name",
		"--embed-author-url", url, "--embed-author-icon", url)
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbedTimestamp(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	args := os.Args[0:1]
	args = append(args, "send", "--embed", "--embed-title", "title", "--embed-timestamp", "2015-12-31T12:00:00.000Z")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbedFields(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.DisableSliceFlagSeparator = true
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	args := os.Args[0:1]
	args = append(args, "send", "--embed",
		"--embed-field", "name,value",
		"--embed-field", "name,value,true",
		"--embed-field", "name,value,true")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithPayload(t *testing.T) {
	is := is.New(t)

	app := cli.NewApp()
	app.DisableSliceFlagSeparator = true
	app.Flags = getFlags()
	app.Commands = []*cli.Command{&command.SendCommand}

	args := os.Args[0:1]
	args = append(args, "send",
		"--payload", "{\"content\":\"content from payload\",\"username\":\"spidee\"}")
	err := app.Run(args)

	is.NoErr(err)
}
