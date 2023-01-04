package cmd

import (
	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/pkg/args"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/urfave/cli/v2"
)

var Version string

func NewApp() *cli.App {
	return app
}

var app *cli.App = &cli.App{
	Name:                      "spidee",
	Usage:                     "Discord webhook CLI",
	Flags:                     flags,
	Commands:                  commands,
	Action:                    action,
	Version:                   Version,
	DisableSliceFlagSeparator: true,
	UseShortOptionHandling:    true,
	HideHelpCommand:           true,
	CustomAppHelpTemplate:     helpTemplate,
}

var commands []*cli.Command = []*cli.Command{}

var flags []cli.Flag = []cli.Flag{
	&cli.IntFlag{
		Name:     args.FlagId,
		Usage:    "webhook id",
		EnvVars:  []string{"SPIDEE_WEBHOOK_ID"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     args.FlagToken,
		Usage:    "webhook token",
		EnvVars:  []string{"SPIDEE_WEBHOOK_TOKEN"},
		Required: true,
	},
}

func action(ctx *cli.Context) error {
	id := discord.WebhookID(ctx.Int(args.FlagId))
	token := ctx.String(args.FlagToken)

	client := webhook.New(id, token)
	context.WrapClient(ctx, client)

	return nil
}

const helpTemplate string = `{{.Name}} {{if .Version}}{{if not .HideVersion}}{{.Version}}{{end}}{{end}}
{{if .Usage}}{{.Usage}}{{end}}

Usage:
{{"\t\t"}}{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .Commands}} command [subcommand] [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Description}}

Description:
{{"\t\t"}}{{.Description | nindent 3 | trim}}{{end}}{{if .VisibleCommands}}

Commands:{{range .VisibleCategories}}{{if .Name}}
{{.Name}}:{{range .VisibleCommands}}
{{join .Names ", "}}{{"\t\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
{{"\t\t"}}{{join .Names ", "}}{{"\t\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

Options:
{{"\t\t"}}{{range $index, $option := .VisibleFlags}}{{if $index}}
{{"\t\t"}}{{end}}{{$option}}{{end}}{{end}}
`
