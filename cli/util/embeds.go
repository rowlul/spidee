package util

import (
	"strconv"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/urfave/cli/v2"
)

func buildEmbedTimestampFromContext(c *cli.Context) (discord.Timestamp, error) {
	var err error

	if c.String("embed-timestamp") == "now" {
		return discord.NowTimestamp(), err
	}

	time, err := time.Parse(time.RFC3339, c.String("embed-timestamp"))
	return discord.NewTimestamp(time), err
}

func buildEmbedFieldsFromContext(c *cli.Context) ([]discord.EmbedField, error) {
	var err error

	var fields []discord.EmbedField
	for _, f := range c.StringSlice("embed-field") {
		field := strings.Split(f, ",")

		var inline bool
		if len(field) > 2 {
			inline, err = strconv.ParseBool(field[2])
		}

		fields = append(fields, discord.EmbedField{Name: field[0], Value: field[1], Inline: inline})
	}

	return fields, err
}

func BuildEmbedsFromContext(c *cli.Context) ([]discord.Embed, error) {
	var embeds []discord.Embed

	timestamp, err := buildEmbedTimestampFromContext(c)
	if len(c.String("embed-timestamp")) > 0 && err != nil {
		return embeds, err
	}

	fields, err := buildEmbedFieldsFromContext(c)
	if err != nil {
		return embeds, err
	}

	embeds = []discord.Embed{{
		Title:       c.String("embed-title"),
		Description: c.String("embed-description"),
		URL:         c.String("embed-url"),
		Timestamp:   timestamp,
		Color:       discord.Color(c.Int("embed-color")),
		Footer:      &discord.EmbedFooter{Text: c.String("embed-footer-text"), Icon: c.String("embed-footer-icon")},
		Image:       &discord.EmbedImage{URL: c.String("embed-image-url")},
		Thumbnail:   &discord.EmbedThumbnail{URL: c.String("embed-thumbnail-url")},
		Video:       &discord.EmbedVideo{URL: c.String("embed-video-url")},
		Provider:    &discord.EmbedProvider{Name: c.String("embed-provider-name"), URL: c.String("embed-provider-url")},
		Author:      &discord.EmbedAuthor{Name: c.String("embed-author-name"), URL: c.String("embed-author-url"), Icon: c.String("embed-author-icon")},
		Fields:      fields},
	}

	return embeds, embeds[0].Validate()
}
