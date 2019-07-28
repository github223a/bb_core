package core

import "log"

func FailOnError(header, message string, err error) {
	if err != nil {
		log.Fatalf("%s %s %s", header, message, err)
	}
}

func LogDebug(header, message string) {
	log.Printf("%s %s", header, message)
}
