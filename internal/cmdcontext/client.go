package cmdcontext

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api/webhook"
	"github.com/urfave/cli/v2"
)

type c struct{}

// WrapClient adds Discord Webhook API client object to cli inner context with a
// unique key of unexported type to be retrieved by calling UnwrapClient.
func WrapClient(ctx *cli.Context, client *webhook.Client) {
	ctx.Context = context.WithValue(ctx.Context, c{}, client)
}

// UnwrapClient returns Discord Webhook API client object from cli inner context,
// or nil if WrapClient had not been previously called.
func UnwrapClient(ctx *cli.Context) *webhook.Client {
	return ctx.Context.Value(c{}).(*webhook.Client)
}
