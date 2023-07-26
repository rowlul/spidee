package cmdcontext

import (
	"strconv"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/internal"
	"github.com/urfave/cli/v2"
)

// AnyEmbedFlag looks through cli.Context flags and returns a bool determining
// whether any flags that contain "embed" are passed.
func AnyEmbedFlag(ctx *cli.Context) bool {
	for _, flag := range ctx.LocalFlagNames() {
		if strings.Contains(flag, "embed") {
			return true
		}
	}

	return false
}

// EmbedTimestamp parses an RFC3339 timestamp from string flag and returns a
// discord.Timestamp. If "now" is passed instead of timestamp, function will
// return a timestamp with current date and time.
func EmbedTimestamp(ctx *cli.Context) (discord.Timestamp, error) {
	timestamp := ctx.String(internal.FlagEmbedTimestamp)

	if timestamp == "now" {
		return discord.NowTimestamp(), nil
	}

	time, err := time.Parse(time.RFC3339, timestamp)
	return discord.NewTimestamp(time), err
}

// EmbedFields loops through string slice flag, extracts 3 values (name,
// value,inline) from each string separated by comma, and returns a slice of
// discord.EmbedField. Last value (inline) is a bool that is parsed via strconv
// that defaults to false when omitted. Returns an error if there are fewer than
// 2 values or inline value is invalid.
func EmbedFields(ctx *cli.Context) ([]discord.EmbedField, error) {
	var fields []discord.EmbedField
	for _, f := range ctx.StringSlice(internal.FlagEmbedField) {
		field := strings.Split(f, ",")

		var inline bool
		if len(field) > 2 {
			result, err := strconv.ParseBool(field[2])
			if err != nil {
				return nil, err
			}

			inline = result
		}

		fields = append(fields, discord.EmbedField{Name: field[0], Value: field[1], Inline: inline})
	}

	return fields, nil
}

// Embeds returns a slice of discord.Embed with a single embed object comprised
// of all available properties. Returns an error if EmbedTimestamp or EmbedFields
// failed, or embed is invalid.
func Embeds(ctx *cli.Context) ([]discord.Embed, error) {
	var embeds []discord.Embed

	title := ctx.String(internal.FlagEmbedTitle)
	description := ctx.String(internal.FlagEmbedDescription)
	url := ctx.String(internal.FlagEmbedURL)
	color := discord.Color(ctx.Int(internal.FlagEmbedColor))
	footer := &discord.EmbedFooter{Text: ctx.String(internal.FlagEmbedFooterText), Icon: ctx.String(internal.FlagEmbedFooterIcon)}
	image := &discord.EmbedImage{URL: ctx.String(internal.FlagEmbedImageURL)}
	thumbnail := &discord.EmbedThumbnail{URL: ctx.String(internal.FlagEmbedThumbnailURL)}
	author := &discord.EmbedAuthor{Name: ctx.String(internal.FlagEmbedAuthorName), URL: ctx.String(internal.FlagEmbedAuthorURL), Icon: ctx.String(internal.FlagEmbedAuthorIcon)}

	timestamp, err := EmbedTimestamp(ctx)
	if ctx.String(internal.FlagEmbedTimestamp) != "" && err != nil {
		return nil, err
	}

	fields, err := EmbedFields(ctx)
	if err != nil {
		return nil, err
	}

	embeds = make([]discord.Embed, 1)
	embeds[0] = discord.Embed{
		Title:       title,
		Description: description,
		URL:         url,
		Timestamp:   timestamp,
		Color:       color,
		Footer:      footer,
		Image:       image,
		Thumbnail:   thumbnail,
		Author:      author,
		Fields:      fields,
	}

	return embeds, embeds[0].Validate()
}
