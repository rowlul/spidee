package cli

import (
	"os"
	"strconv"
	"testing"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/matryer/is"
	"github.com/rowlul/spidee/cli/command"
	"github.com/urfave/cli/v2"
)

func getWebhookClient() *webhook.Client {
	id, err := strconv.Atoi(os.Getenv("SPIDEE_WEBHOOK_ID"))
	if err != nil {
		panic("no or invalid id supplied")
	}

	token := os.Getenv("SPIDEE_WEBHOOK_TOKEN")
	if token == "" {
		panic("no token supplied")
	}

	return webhook.New(discord.WebhookID(id), token)
}

func TestSendCommandWithContent(t *testing.T) {
	is := is.New(t)
	app := cli.NewApp()
	client := getWebhookClient()
	command := command.SendCommand(*client)

	app.Commands = []*cli.Command{&command}

	args := os.Args[0:1]
	args = append(args, "send", "--content", "content")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithFile(t *testing.T) {
	is := is.New(t)
	app := cli.NewApp()
	client := getWebhookClient()
	command := command.SendCommand(*client)

	app.Commands = []*cli.Command{&command}

	args := os.Args[0:1]
	args = append(args, "send", "--file", "app_test.go")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbed(t *testing.T) {
	is := is.New(t)
	app := cli.NewApp()
	client := getWebhookClient()
	command := command.SendCommand(*client)

	app.Commands = []*cli.Command{&command}

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
	client := getWebhookClient()
	command := command.SendCommand(*client)

	app.Commands = []*cli.Command{&command}

	args := os.Args[0:1]
	args = append(args, "send", "--embed", "--embed-title", "title", "--embed-timestamp", "2015-12-31T12:00:00.000Z")
	err := app.Run(args)

	is.NoErr(err)
}

func TestSendCommandWithEmbedFields(t *testing.T) {
	is := is.New(t)
	app := cli.NewApp()
	app.DisableSliceFlagSeparator = true
	client := getWebhookClient()
	command := command.SendCommand(*client)

	app.Commands = []*cli.Command{&command}

	args := os.Args[0:1]
	args = append(args, "send", "--embed",
		"--embed-field", "name,value",
		"--embed-field", "name,value,true",
		"--embed-field", "name,value,true")
	err := app.Run(args)

	is.NoErr(err)
}
