package cmdcontext

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/urfave/cli/v2"
)

func Image(c *cli.Context) (*api.Image, error) {
	b, err := os.ReadFile(c.String("avatar"))
	if err != nil {
		return nil, err
	}

	var b64 string
	switch http.DetectContentType(b) {
	case "image/jpeg":
		b64 += "data:image/jpeg;base64,"
	case "image/png":
		b64 += "data:image/png;base64,"
	case "image/gif":
		b64 += "data:image/gif;base64,"
	default:
		return nil, api.ErrInvalidImageCT
	}

	b64 += base64.StdEncoding.EncodeToString(b)

	image, err := api.DecodeImage([]byte(b64))
	return image, err
}
