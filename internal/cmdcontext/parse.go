package cmdcontext

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rowlul/spidee/internal/vt"
	"github.com/urfave/cli/v2"
)

func EnsureFlags(ctx *cli.Context) error {
	if ctx.NumFlags() == 0 && !vt.IsStdin() {
		return errors.New("no flags supplied")
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
