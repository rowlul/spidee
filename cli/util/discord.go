package util

import (
	"encoding/json"

	"github.com/diamondburned/arikawa/v3/discord"
)

func StringifyMessage(message *discord.Message, format bool) (string, error) {
	if message == nil {
		return "{}", nil
	}

	var (
		marshalled []byte
		err        error
	)

	if format {
		marshalled, err = json.MarshalIndent(message, "", "    ")
	} else {
		marshalled, err = json.Marshal(message)
	}

	return string(marshalled), err
}
