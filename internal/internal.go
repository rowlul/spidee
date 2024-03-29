package internal

import "github.com/urfave/cli/v2"

const (
	CommandSend   = "send"
	CommandEdit   = "edit"
	CommandDelete = "delete"
	CommandGet    = "get"
	CommandSelf   = "self"
	CommandModify = "modify"
)

const (
	FlagId      = "id"
	FlagToken   = "token"
	FlagVersion = "version"
)

const (
	FlagContent   = "content"
	FlagUsername  = "username"
	FlagAvatarURL = "avatar-url"
	FlagAvatar    = "avatar"
	FlagTTS       = "tts"
	FlagFile      = "file"
	FlagPayload   = "payload"
	FlagWait      = "wait"
	FlagJSON      = "json"
	FlagNoRedact  = "no-redact"
	FlagThreadId  = "thread-id"
)

const (
	FlagEmbedTitle        = "embed-title"
	FlagEmbedDescription  = "embed-description"
	FlagEmbedURL          = "embed-url"
	FlagEmbedTimestamp    = "embed-timestamp"
	FlagEmbedColor        = "embed-color"
	FlagEmbedFooterText   = "embed-footer-text"
	FlagEmbedFooterIcon   = "embed-footer-icon"
	FlagEmbedImageURL     = "embed-image-url"
	FlagEmbedThumbnailURL = "embed-thumbnail-url"
	FlagEmbedAuthorName   = "embed-author-name"
	FlagEmbedAuthorURL    = "embed-author-url"
	FlagEmbedAuthorIcon   = "embed-author-icon"
	FlagEmbedField        = "embed-field"
)

var EmbedFlags = []cli.Flag{
	&cli.StringFlag{Name: FlagEmbedTitle, Usage: "embed title", Aliases: []string{"et"}},
	&cli.StringFlag{Name: FlagEmbedDescription, Usage: "embed description", Aliases: []string{"ed"}},
	&cli.StringFlag{Name: FlagEmbedURL, Usage: "embed url", Aliases: []string{"eu"}},
	&cli.StringFlag{Name: FlagEmbedTimestamp, Usage: "embed timestamp (now|RFC3339 timestamp)", Aliases: []string{"eT"}},
	&cli.IntFlag{Name: FlagEmbedColor, Usage: "embed color (hex)", Aliases: []string{"ec"}},
	&cli.StringFlag{Name: FlagEmbedFooterText, Usage: "embed footer text", Aliases: []string{"eft"}},
	&cli.StringFlag{Name: FlagEmbedFooterIcon, Usage: "embed footer icon", Aliases: []string{"efi"}},
	&cli.StringFlag{Name: FlagEmbedImageURL, Usage: "embed image url", Aliases: []string{"eiu"}},
	&cli.StringFlag{Name: FlagEmbedThumbnailURL, Usage: "embed thumbnail url", Aliases: []string{"etu"}},
	&cli.StringFlag{Name: FlagEmbedAuthorName, Usage: "embed author name", Aliases: []string{"ean"}},
	&cli.StringFlag{Name: FlagEmbedAuthorURL, Usage: "embed author url", Aliases: []string{"eau"}},
	&cli.StringFlag{Name: FlagEmbedAuthorIcon, Usage: "embed author icon", Aliases: []string{"eai"}},
	&cli.StringSliceFlag{Name: FlagEmbedField, Usage: "embed field (name,value,inline)", Aliases: []string{"ef"}},
}
