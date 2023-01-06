package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/rowlul/spidee/pkg/vt"
	"github.com/urfave/cli/v2"
)

func NewEditCommand() *cli.Command {
	cmd := &cli.Command{
		Name:      pkg.CommandEdit,
		Usage:     "Edit webhook message",
		ArgsUsage: "<message id>",
		Before:    beforeEdit,
		Action:    actionEdit,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: pkg.FlagContent, Usage: "plain text", Aliases: []string{"c"}},
			&cli.StringSliceFlag{Name: pkg.FlagFile, Usage: "webhook attachment", Aliases: []string{"f"}, TakesFile: true},
			&cli.StringFlag{Name: pkg.FlagPayload, Usage: "raw json payload"},
			&cli.BoolFlag{Name: pkg.FlagJSON, Usage: "return JSON message object"},
		},
	}

	cmd.Flags = append(cmd.Flags, pkg.EmbedFlags...)

	return cmd
}

func beforeEdit(ctx *cli.Context) error {
	if err := context.EnsureFlags(ctx); err != nil {
		return err
	}

	return nil
}

func actionEdit(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	arg, err := context.Uint64Arg(ctx)
	if err != nil {
		return err
	}

	id := discord.MessageID(arg)

	var data webhook.EditMessageData
	var (
		content = ctx.String(pkg.FlagContent)
		payload = ctx.String(pkg.FlagPayload)
	)

	if vt.IsStdin() {
		input := strings.Join(vt.ReadStdin(), "\n")
		if json.Valid([]byte(input)) {
			payload = input
		} else {
			content = input
		}
	}

	if payload == "" {
		files, err := context.Files(ctx)
		if err != nil {
			return err
		}

		data = webhook.EditMessageData{
			Content: option.NewNullableString(content),
			Files:   files,
		}

		if context.AnyEmbedFlag(ctx) {
			embeds, err := context.Embeds(ctx)
			if err != nil {
				return err
			}

			data.Embeds = &embeds
		}
	} else {
		json.Unmarshal([]byte(payload), &data)
	}

	message, err := client.EditMessage(id, data)
	if err != nil {
		return err
	}

	if ctx.Bool(pkg.FlagJSON) {
		json, err := json.Marshal(message)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
