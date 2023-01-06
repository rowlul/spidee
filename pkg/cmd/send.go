package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/rowlul/spidee/pkg/vt"
	"github.com/urfave/cli/v2"
)

func NewSendCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   pkg.CommandSend,
		Usage:  "Send webhook message",
		Before: beforeSend,
		Action: actionSend,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: pkg.FlagContent, Usage: "plain text", Aliases: []string{"c"}},
			&cli.StringFlag{Name: pkg.FlagUsername, Usage: "webhook username", Aliases: []string{"u"}},
			&cli.StringFlag{Name: pkg.FlagAvatarURL, Usage: "webhook avatar url", Aliases: []string{"a"}},
			&cli.BoolFlag{Name: pkg.FlagTTS, Usage: "narrate message"},
			&cli.StringSliceFlag{Name: pkg.FlagFile, Usage: "webhook attachment", Aliases: []string{"f"}, TakesFile: true},
			&cli.StringFlag{Name: pkg.FlagPayload, Usage: "raw json payload"},
			&cli.BoolFlag{Name: pkg.FlagWait, Usage: "wait for message to be created", Aliases: []string{"w"}},
			&cli.BoolFlag{Name: pkg.FlagJSON, Usage: "return JSON message object"},
		},
	}

	cmd.Flags = append(cmd.Flags, pkg.EmbedFlags...)

	return cmd
}

func beforeSend(ctx *cli.Context) error {
	if err := context.EnsureFlags(ctx); err != nil {
		return err
	}

	if !ctx.Bool(pkg.FlagWait) && ctx.Bool(pkg.FlagJSON) {
		return errors.New("must wait for message before output (use --wait)")
	}

	return nil
}

func actionSend(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	var data webhook.ExecuteData
	var (
		content = ctx.String(pkg.FlagContent)
		payload = ctx.String(pkg.FlagPayload)
		wait    = ctx.Bool(pkg.FlagWait)
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

		var (
			username  = ctx.String(pkg.FlagUsername)
			avatarUrl = ctx.String(pkg.FlagAvatarURL)
			tts       = ctx.Bool(pkg.FlagTTS)
		)

		data = webhook.ExecuteData{
			Content:   content,
			Username:  username,
			AvatarURL: avatarUrl,
			TTS:       tts,
			Files:     files,
		}

		if context.AnyEmbedFlag(ctx) {
			embeds, err := context.Embeds(ctx)
			if err != nil {
				return err
			}

			data.Embeds = embeds
		}
	} else {
		json.Unmarshal([]byte(payload), &data)
	}

	if !wait {
		if err := client.Execute(data); err != nil {
			return err
		}
	} else {
		message, err := client.ExecuteAndWait(data)
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
	}

	return nil
}
