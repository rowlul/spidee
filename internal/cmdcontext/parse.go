package cmdcontext

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// EnsureFlags looks through cli.Context flags and ensures that there are any flags
// passed to context excluding flags specified in ignoredFlags. If no ignored flags
// are specified, this function will return an error if there are no flags whatsoever.
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

// Uint64Arg parses first arg in cli.Context and returns a uint64 value. If an
// error occurred, 0 and error will be returned respectively.
func Uint64Arg(ctx *cli.Context) (uint64, error) {
	arg := ctx.Args().First()

	parsed, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("could not parse argument as uint64: %q", arg)
	}

	return parsed, nil
}
