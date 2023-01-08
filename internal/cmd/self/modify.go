package self

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/rowlul/spidee/internal"
	"github.com/rowlul/spidee/internal/cmdcontext"
	"github.com/rowlul/spidee/internal/vt"
	"github.com/urfave/cli/v2"
)

func NewModifyCommand() *cli.Command {
	cmd := &cli.Command{
		Name:         internal.CommandModify,
		Usage:        "Modify webhook",
		Before:       beforeEdit,
		Action:       actionEdit,
		OnUsageError: usageError,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: internal.FlagUsername, Usage: "webhook username", Aliases: []string{"u"}},
			&cli.StringFlag{Name: internal.FlagAvatar, Usage: "path to webhook avatar image", Aliases: []string{"a"}, TakesFile: true},
			&cli.StringFlag{Name: internal.FlagPayload, Usage: "raw json payload"},
			&cli.BoolFlag{Name: internal.FlagJSON, Usage: "return JSON message object"},
		},
	}

	return cmd
}

func beforeEdit(ctx *cli.Context) error {
	ignoredFlags := []string{
		internal.FlagJSON,
	}

	if err := cmdcontext.EnsureFlags(ctx, ignoredFlags...); err != nil {
		cli.ShowSubcommandHelp(ctx)
		return errors.New("no username or avatar supplied")
	}

	if ctx.String(internal.FlagUsername) == "" && ctx.String(internal.FlagAvatar) != "" {
		cli.ShowSubcommandHelp(ctx)
		return errors.New("username must be supplied along with avatar")
	}

	return nil
}

func actionEdit(ctx *cli.Context) error {
	client := cmdcontext.UnwrapClient(ctx)

	var data api.ModifyWebhookData

	payload := ctx.String(internal.FlagPayload)
	if vt.IsStdin() {
		input := strings.Join(vt.ReadStdin(), "\n")
		if json.Valid([]byte(input)) {
			payload = input
		}
	}

	if payload == "" {
		avatar, err := cmdcontext.Image(ctx)
		if ctx.String(internal.FlagAvatar) != "" && err != nil {
			return err
		}

		username := ctx.String(internal.FlagUsername)

		data = api.ModifyWebhookData{
			Name:   &username,
			Avatar: avatar,
		}
	} else {
		json.Unmarshal([]byte(payload), &data)
	}

	webhook, err := client.Modify(data)
	if err != nil {
		return err
	}

	if ctx.Bool(internal.FlagJSON) {
		json, err := json.Marshal(webhook)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
