package cmd

import (
	"fmt"
	"os"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/cmd/self"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/rowlul/spidee/pkg/vt"
	"github.com/urfave/cli/v2"
)

var Version string

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "spidee",
		Usage: "Discord webhook CLI",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:     pkg.FlagId,
				Usage:    "webhook id",
				EnvVars:  []string{"SPIDEE_WEBHOOK_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     pkg.FlagToken,
				Usage:    "webhook token",
				EnvVars:  []string{"SPIDEE_WEBHOOK_TOKEN"},
				Required: true,
			},
		},
		Commands: []*cli.Command{
			NewSendCommand(),
			NewEditCommand(),
			NewDeleteCommand(),
			NewGetCommand(),
			self.NewSelfCommand(),
		},
		Before:                    before,
		CommandNotFound:           cmdNotFound,
		Version:                   Version,
		DisableSliceFlagSeparator: true,
		UseShortOptionHandling:    true,
		HideHelpCommand:           true,
	}

	if vt.IsStdin() {
		app.DefaultCommand = pkg.CommandSend
	}

	return app
}

func before(ctx *cli.Context) error {
	id := discord.WebhookID(ctx.Uint64(pkg.FlagId))
	token := ctx.String(pkg.FlagToken)

	client := webhook.New(id, token)
	context.WrapClient(ctx, client)

	return nil
}

func cmdNotFound(ctx *cli.Context, s string) {
	cli.ShowAppHelp(ctx)
	fmt.Fprintln(os.Stderr, "no matching command:", s)
	os.Exit(1)
}
