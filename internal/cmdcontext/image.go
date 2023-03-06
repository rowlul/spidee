package cmdcontext

import (
	"encoding/base64"
	"net/http"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/rowlul/spidee/internal"
	"github.com/urfave/cli/v2"
)

// Image reads file from path specified in corresponding context flag, encodes
// bytes to Base64, and returns an Image. In case image format is not jpeg, png, or
// gif, an error will be returned.
// https://discord.com/developers/docs/reference#image-data
func Image(c *cli.Context) (*api.Image, error) {
	b, err := os.ReadFile(c.String(internal.FlagAvatar))
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
