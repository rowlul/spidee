package util

import (
	"log"
	"strings"
)

func HandleError(err error) {
	log.SetFlags(0)

	e := err.Error()
	if strings.Contains(e, "strconv.Atoi") {
		log.Fatalln(e, "\nRequired argument not set or integer")
	} else if strings.Contains(e, "strconv.ParseBool") {
		log.Fatalln(e, "\nArgument must be bool")
	} else {
		log.Fatalln(e)
	}
}
