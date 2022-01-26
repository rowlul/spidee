package util

import (
	"reflect"
	"testing"
	"time"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/matryer/is"
	"github.com/urfave/cli/v2"
)

func TestBuildEmbedTimestampFromContext(t *testing.T) {
	is := is.New(t)

	timeStr := "2015-12-31T12:00:00.000Z"
	args := []string{"spidee", "--embed-timestamp", timeStr}
	app := &cli.App{
		Flags: []cli.Flag{&cli.StringFlag{Name: "embed-timestamp"}},
		Action: func(c *cli.Context) error {
			timestamp, err := buildEmbedTimestampFromContext(c)
			is.NoErr(err)

			a := timestamp.Time()
			b, err := time.Parse(time.RFC3339, timeStr)
			is.NoErr(err)

			is.Equal(a, b)

			return nil
		},
	}

	app.Run(args)
}

func TestBuildEmbedFieldsFromContext(t *testing.T) {
	is := is.New(t)

	args := []string{"spidee", "--embed-field", "name,value,true"}
	app := &cli.App{
		Flags: []cli.Flag{&cli.StringSliceFlag{Name: "embed-field"}},
		Action: func(c *cli.Context) error {
			fields, err := buildEmbedFieldsFromContext(c)
			is.NoErr(err)

			a := fields[0]
			b := discord.EmbedField{Name: "name", Value: "value", Inline: true}

			is.Equal(a, b)

			return nil
		},
	}

	app.Run(args)
}

func TestBuildEmbedsFromContext(t *testing.T) {
	is := is.New(t)

	url := "https://go.dev/"
	args := []string{
		"spidee",
		"--embed-title", "title", "--embed-description", "description", "--embed-url", url, "--embed-color", "990000",
		"--embed-footer-text", "text", "--embed-footer-icon", url, "--embed-image-url", url, "--embed-thumbnail-url", url,
		"--embed-video-url", url, "--embed-provider-name", "name", "--embed-provider-url", url, "--embed-author-name", "name",
		"--embed-author-url", url, "--embed-author-icon", url,
		"--embed-timestamp", "2015-12-31T12:00:00.000Z",
		"--embed-field", "name,value,true"}

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "embed-title"},
			&cli.StringFlag{Name: "embed-description"},
			&cli.StringFlag{Name: "embed-url"},
			&cli.StringFlag{Name: "embed-timestamp"},
			&cli.IntFlag{Name: "embed-color"},
			&cli.StringFlag{Name: "embed-footer-text"},
			&cli.StringFlag{Name: "embed-footer-icon"},
			&cli.StringFlag{Name: "embed-image-url"},
			&cli.StringFlag{Name: "embed-thumbnail-url"},
			&cli.StringFlag{Name: "embed-video-url"},
			&cli.StringFlag{Name: "embed-provider-name"},
			&cli.StringFlag{Name: "embed-provider-url"},
			&cli.StringFlag{Name: "embed-author-name"},
			&cli.StringFlag{Name: "embed-author-url"},
			&cli.StringFlag{Name: "embed-author-icon"},
			&cli.StringSliceFlag{Name: "embed-field"},
		},
		Action: func(c *cli.Context) error {
			a, err := BuildEmbedsFromContext(c)
			is.NoErr(err)

			b := []discord.Embed{{
				Type:        discord.EmbedType("rich"),
				Title:       "title",
				Description: "description",
				URL:         url,
				Color:       discord.Color(990000),
				Timestamp:   discord.NewTimestamp(time.Date(2015, time.December, 31, 12, 00, 00, 00, time.UTC)),
				Footer:      &discord.EmbedFooter{Text: "text", Icon: url},
				Image:       &discord.EmbedImage{URL: url},
				Video:       &discord.EmbedVideo{URL: url},
				Thumbnail:   &discord.EmbedThumbnail{URL: url},
				Provider:    &discord.EmbedProvider{Name: "name", URL: url},
				Author:      &discord.EmbedAuthor{Name: "name", URL: url, Icon: url},
				Fields: []discord.EmbedField{
					{
						Name:   "name",
						Value:  "value",
						Inline: true,
					},
				},
			}}

			aV := reflect.ValueOf(a[0])
			bV := reflect.ValueOf(b[0])
			for i := 0; i < aV.NumField(); i++ {
				for j := 0; j < bV.NumField(); j++ {
					is.Equal(aV.Field(i).Interface(), bV.Field(i).Interface())
				}
			}

			return nil
		},
	}

	app.Run(args)
}
