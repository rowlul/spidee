package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/context"
	"github.com/rowlul/spidee/internal/vt"
	"github.com/urfave/cli/v2"
)

func NewSendCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   internal.CommandSend,
		Usage:  "Send webhook message",
		Before: beforeSend,
		Action: actionSend,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: internal.FlagContent, Usage: "plain text", Aliases: []string{"c"}},
			&cli.StringFlag{Name: internal.FlagUsername, Usage: "webhook username", Aliases: []string{"u"}},
			&cli.StringFlag{Name: internal.FlagAvatarURL, Usage: "webhook avatar url", Aliases: []string{"a"}},
			&cli.BoolFlag{Name: internal.FlagTTS, Usage: "narrate message"},
			&cli.StringSliceFlag{Name: internal.FlagFile, Usage: "webhook attachment", Aliases: []string{"f"}, TakesFile: true},
			&cli.StringFlag{Name: internal.FlagPayload, Usage: "raw json payload"},
			&cli.BoolFlag{Name: internal.FlagWait, Usage: "wait for message to be created", Aliases: []string{"w"}},
			&cli.BoolFlag{Name: internal.FlagJSON, Usage: "return JSON message object"},
		},
	}

	cmd.Flags = append(cmd.Flags, internal.EmbedFlags...)

	return cmd
}

func beforeSend(ctx *cli.Context) error {
	if err := context.EnsureFlags(ctx); err != nil {
		return err
	}

	if !ctx.Bool(internal.FlagWait) && ctx.Bool(internal.FlagJSON) {
		return errors.New("must wait for message before output (use --wait)")
	}

	return nil
}

func actionSend(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	var data webhook.ExecuteData
	var (
		content = ctx.String(internal.FlagContent)
		payload = ctx.String(internal.FlagPayload)
		wait    = ctx.Bool(internal.FlagWait)
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
			username  = ctx.String(internal.FlagUsername)
			avatarUrl = ctx.String(internal.FlagAvatarURL)
			tts       = ctx.Bool(internal.FlagTTS)
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

		if ctx.Bool(internal.FlagJSON) {
			json, err := json.Marshal(message)
			if err != nil {
				return err
			}

			fmt.Println(string(json))
		}
	}

	return nil
}
