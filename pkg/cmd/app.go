package cmd

import (
	"github.com/urfave/cli/v2"
)

var Version string

func NewApp() *cli.App {
	app := &cli.App{
		Name:     "spidee",
		Usage:    "Discord webhook CLI",
		Commands: []*cli.Command{},
		Version:  Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "id",
				Usage:    "webhook id",
				EnvVars:  []string{"SPIDEE_WEBHOOK_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "token",
				Usage:    "webhook token",
				EnvVars:  []string{"SPIDEE_WEBHOOK_TOKEN"},
				Required: true,
			},
		},
		DisableSliceFlagSeparator: true,
		UseShortOptionHandling:    true,
		HideHelpCommand:           true,
		CustomAppHelpTemplate:     helpTemplate,
	}

	return app
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
