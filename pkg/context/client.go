package context

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/urfave/cli/v2"
)

type c struct{}

func WrapClient(ctx *cli.Context, client *webhook.Client) {
	ctx.Context = context.WithValue(ctx.Context, c{}, client)
}

func UnwrapClient(ctx *cli.Context) *webhook.Client {
	return ctx.Context.Value(c{}).(*webhook.Client)
}
