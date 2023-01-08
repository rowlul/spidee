package cmdcontext

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

func EnsureFlags(ctx *cli.Context, ignoredFlags ...string) error {
	flags := ctx.LocalFlagNames()

	for _, f := range ignoredFlags {
		for i, v := range flags {
			if v == f {
				flags = append(flags[:i], flags[i+1:]...)
				break
			}
		}
	}

	if len(flags) == 0 {
		return errors.New("no valid flags supplied")
	}

	return nil
}

func Uint64Arg(ctx *cli.Context) (uint64, error) {
	arg := ctx.Args().First()

	parsed, err := strconv.ParseUint(string(arg), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse argument as uint64: %q", arg)
	}

	return parsed, nil
}
