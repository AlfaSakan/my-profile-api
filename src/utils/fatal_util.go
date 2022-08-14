package utils

import "log"

func Fatal(err error, message string) {
	if err != nil {
		log.Fatalf("%s => %s", message, err)
	}
}
