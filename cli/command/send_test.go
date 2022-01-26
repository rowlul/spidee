package command_test

import (
	"testing"

	"github.com/rowlul/spidee/cli"
	"github.com/matryer/is"
)

func TestSendCommandWithContent(t *testing.T) {
	is := is.New(t)

	args := []string{"spidee", "send", "--content", "content"}
	err := cli.App.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithFile(t *testing.T) {
	is := is.New(t)

	args := []string{"spidee", "send", "--file", "send_test.go"}
	err := cli.App.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbed(t *testing.T) {
	is := is.New(t)

	url := "https://go.dev/"
	args := []string{
		"spidee", "send", "--embed",
		"--embed-title", "title", "--embed-description", "description", "--embed-url", url, "--embed-color", "990000",
		"--embed-footer-text", "text", "--embed-footer-icon", url, "--embed-image-url", url, "--embed-thumbnail-url", url,
		"--embed-video-url", url, "--embed-provider-name", "name", "--embed-provider-url", url, "--embed-author-name", "name",
		"--embed-author-url", url, "--embed-author-icon", url}
	err := cli.App.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbedTimestamp(t *testing.T) {
	is := is.New(t)

	args := []string{"spidee", "send", "--embed", "--embed-title", "title", "--embed-timestamp", "2015-12-31T12:00:00.000Z"}
	err := cli.App.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbedFields(t *testing.T) {
	is := is.New(t)

	args := []string{
		"spidee", "send", "--embed",
		"--embed-field", "name,value",
		"--embed-field", "name,value,true",
		"--embed-field", "name,value,true",
	}
	err := cli.App.Run(args)

	is.NoErr(err)
}
