package util

import (
	"encoding/json"
)

func StringifyObject(obj interface{}, format bool) (string, error) {
	if obj == nil {
		return "{}", nil
	}

	var (
		marshalled []byte
		err        error
	)

	if format {
		marshalled, err = json.MarshalIndent(obj, "", "    ")
	} else {
		marshalled, err = json.Marshal(obj)
	}

	return string(marshalled), err
}
