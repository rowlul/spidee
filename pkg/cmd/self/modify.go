package self

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/rowlul/spidee/pkg"
	"github.com/rowlul/spidee/pkg/context"
	"github.com/rowlul/spidee/pkg/vt"
	"github.com/urfave/cli/v2"
)

func NewModifyCommand() *cli.Command {
	cmd := &cli.Command{
		Name:   pkg.CommandModify,
		Usage:  "Modify webhook",
		Before: beforeEdit,
		Action: actionEdit,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: pkg.FlagUsername, Usage: "webhook username", Aliases: []string{"u"}},
			&cli.StringFlag{Name: pkg.FlagAvatar, Usage: "path to webhook avatar image", Aliases: []string{"a"}, TakesFile: true},
			&cli.StringFlag{Name: pkg.FlagPayload, Usage: "raw json payload"},
			&cli.BoolFlag{Name: pkg.FlagJSON, Usage: "return JSON message object"},
		},
	}

	return cmd
}

func beforeEdit(ctx *cli.Context) error {
	if err := context.EnsureFlags(ctx); err != nil {
		return err
	}

	if ctx.String(pkg.FlagUsername) == "" && ctx.String(pkg.FlagAvatar) != "" {
		return errors.New("username must be supplied along with avatar")
	}

	return nil
}

func actionEdit(ctx *cli.Context) error {
	client := context.UnwrapClient(ctx)

	var data api.ModifyWebhookData

	payload := ctx.String(pkg.FlagPayload)
	if vt.IsStdin() {
		input := strings.Join(vt.ReadStdin(), "\n")
		if json.Valid([]byte(input)) {
			payload = input
		}
	}

	if payload == "" {
		avatar, err := context.Image(ctx)
		if ctx.String(pkg.FlagAvatar) != "" && err != nil {
			return err
		}

		username := ctx.String(pkg.FlagUsername)

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

	if ctx.Bool(pkg.FlagJSON) {
		json, err := json.Marshal(webhook)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}
