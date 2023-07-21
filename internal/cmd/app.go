package cmd

import (
	"fmt"
	"os"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmd/self"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/rowlul/spidee/internal/vt"
	"github.com/urfave/cli/v2"
)

// Version specified via ldflags
var Version string

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "spidee",
		Usage: "Discord webhook CLI",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:     internal.FlagId,
				Usage:    "webhook id",
				EnvVars:  []string{"SPIDEE_WEBHOOK_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     internal.FlagToken,
				Usage:    "webhook token",
				EnvVars:  []string{"SPIDEE_WEBHOOK_TOKEN"},
				Required: true,
			},
			&cli.BoolFlag{Name: internal.FlagVersion, Usage: "print name and version", Aliases: []string{"v"}},
		},
		Commands: []*cli.Command{
			NewSendCommand(),
			NewEditCommand(),
			NewDeleteCommand(),
			NewGetCommand(),
			self.NewSelfCommand(),
		},
		Action:                    action,
		Before:                    before,
		CommandNotFound:           cmdNotFound,
		OnUsageError:              usageError,
		Version:                   Version,
		DisableSliceFlagSeparator: true,
		UseShortOptionHandling:    true,
		HideHelpCommand:           true,
	}

	if vt.IsStdin() {
		app.DefaultCommand = internal.CommandSend
	}

	return app
}

func action(ctx *cli.Context) error {
	if ctx.Bool(internal.FlagVersion) {
		fmt.Print(ctx.App.Name, "\n", ctx.App.Version)
		os.Exit(0)
	}

	// since we're overriding default root command behavior, we show help
	// and exit by default if no command passed
	cli.ShowAppHelpAndExit(ctx, 0)

	return nil
}

func before(ctx *cli.Context) error {
	id := discord.WebhookID(ctx.Uint64(internal.FlagId))
	token := ctx.String(internal.FlagToken)

	client := webhook.New(id, token)
	cmdcontext.WrapClient(ctx, client)

	return nil
}

func cmdNotFound(ctx *cli.Context, s string) {
	cli.ShowSubcommandHelp(ctx)
	fmt.Fprintln(os.Stderr, "no matching command:", s)
	os.Exit(1)
}

func usageError(ctx *cli.Context, err error, isSubcommand bool) error {
	if isSubcommand {
		cli.ShowSubcommandHelp(ctx)
	} else {
		cli.ShowAppHelp(ctx)
	}
	return err
}
