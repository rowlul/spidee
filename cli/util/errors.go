package util

import (
	"fmt"
	"strings"
)

func FormatError(err error) string {
	e := err.Error()
	return fmt.Sprint("error: ", e)
}

func FormatFileError(err error) string {
	e := err.Error()
	return fmt.Sprint("file error: ", e)
}

func FormatEmbedError(err error) string {
	e := err.Error()
	if strings.Contains(e, "strconv.ParseBool") {
		e = "field inline must be bool"
	}
	return fmt.Sprint("embed error: ", e)
}
