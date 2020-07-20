package main

import (
	"log"

	"github.com/shellucas/go-commander/flags"
)

func main() {
	b := new(bool)

	s := flags.StringRequired("c", "configuration", "The configuration name of the application")
	d := flags.String("d", "debugmessage", "This is the default debug message", "The debug message that is sent to a person")
	flags.BoolVar(b, "l", "logging", false, "Whether or not the application logs errors")

	flags.Parse()

	log.Printf("Configuration: %s\n", *s)
	log.Printf("Debug Message: %s\n", *d)
	log.Printf("Logging: %t\n", *b)
}
