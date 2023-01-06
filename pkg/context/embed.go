package context

import (
	"strconv"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/rowlul/spidee/pkg"
	"github.com/urfave/cli/v2"
)

func AnyEmbedFlag(ctx *cli.Context) bool {
	for _, flag := range ctx.LocalFlagNames() {
		if strings.Contains(flag, "embed") {
			return true
		}
	}

	return false
}

func EmbedTimestamp(ctx *cli.Context) (discord.Timestamp, error) {
	timestamp := ctx.String(pkg.FlagEmbedTimestamp)

	if timestamp == "now" {
		return discord.NowTimestamp(), nil
	}

	time, err := time.Parse(time.RFC3339, timestamp)
	return discord.NewTimestamp(time), err
}

func EmbedFields(ctx *cli.Context) ([]discord.EmbedField, error) {
	var fields []discord.EmbedField
	for _, f := range ctx.StringSlice(pkg.FlagEmbedField) {
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

func Embeds(ctx *cli.Context) ([]discord.Embed, error) {
	var embeds []discord.Embed

	title := ctx.String(pkg.FlagEmbedTitle)
	description := ctx.String(pkg.FlagEmbedDescription)
	url := ctx.String(pkg.FlagEmbedURL)
	color := discord.Color(ctx.Int(pkg.FlagEmbedColor))
	footer := &discord.EmbedFooter{Text: ctx.String(pkg.FlagEmbedFooterText), Icon: ctx.String(pkg.FlagEmbedFooterIcon)}
	image := &discord.EmbedImage{URL: ctx.String(pkg.FlagEmbedImageURL)}
	thumbnail := &discord.EmbedThumbnail{URL: ctx.String(pkg.FlagEmbedThumbnailURL)}
	video := &discord.EmbedVideo{URL: ctx.String(pkg.FlagEmbedVideoURL)}
	provider := &discord.EmbedProvider{Name: ctx.String(pkg.FlagEmbedProviderName), URL: ctx.String(pkg.FlagEmbedProviderURL)}
	author := &discord.EmbedAuthor{Name: ctx.String(pkg.FlagEmbedAuthorName), URL: ctx.String(pkg.FlagEmbedAuthorURL), Icon: ctx.String(pkg.FlagEmbedAuthorIcon)}

	timestamp, err := EmbedTimestamp(ctx)
	if ctx.String(pkg.FlagEmbedTimestamp) != "" && err != nil {
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
		Video:       video,
		Provider:    provider,
		Author:      author,
		Fields:      fields,
	}

	return embeds, embeds[0].Validate()
}
