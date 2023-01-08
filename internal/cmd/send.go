package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/rowlul/spidee/internal/vt"
	"github.com/urfave/cli/v2"
)

func NewSendCommand() *cli.Command {
	cmd := &cli.Command{
		Name:         internal.CommandSend,
		Usage:        "Send webhook message",
		ArgsUsage:    "[content|payload]",
		Before:       beforeSend,
		Action:       actionSend,
		OnUsageError: usageError,
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
		HideHelpCommand: true,
	}

	cmd.Flags = append(cmd.Flags, internal.EmbedFlags...)

	return cmd
}

func beforeSend(ctx *cli.Context) error {
	if ctx.Args().First() != "" {
		input := ctx.Args().First()
		if json.Valid([]byte(input)) {
			ctx.Set(internal.FlagPayload, input)
		} else {
			ctx.Set(internal.FlagContent, input)
		}
	}

	if vt.IsStdin() {
		input := strings.Join(vt.ReadStdin(), "\n")
		if json.Valid([]byte(input)) {
			ctx.Set(internal.FlagPayload, input)
		} else {
			ctx.Set(internal.FlagContent, input)
		}
	}

	ignoredFlags := []string{
		internal.FlagUsername, internal.FlagAvatarURL, internal.FlagTTS, internal.FlagWait, internal.FlagJSON,
		internal.FlagEmbedURL, internal.FlagEmbedColor, internal.FlagEmbedTimestamp, internal.FlagEmbedAuthorURL,
		"u", "a", "w", "eu", "ec", "et", "eau",
	}

	if err := cmdcontext.EnsureFlags(ctx, ignoredFlags...); err != nil {
		cli.ShowSubcommandHelp(ctx)
		return errors.New("no content, file, or embed supplied")
	}

	if !ctx.Bool(internal.FlagWait) && ctx.Bool(internal.FlagJSON) {
		cli.ShowSubcommandHelp(ctx)
		return errors.New("must wait for message before output (use --wait)")
	}

	return nil
}

func actionSend(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)

	var data webhook.ExecuteData
	var (
		content = ctx.String(internal.FlagContent)
		payload = ctx.String(internal.FlagPayload)
		wait    = ctx.Bool(internal.FlagWait)
	)

	if payload == "" {
		files, err := cmdcontext.Files(ctx)
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

		if cmdcontext.AnyEmbedFlag(ctx) {
			embeds, err := cmdcontext.Embeds(ctx)
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
