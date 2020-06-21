package main

import (
	"log"

	flags "github.com/BlackwonderTF/go-commander"
)

func main() {
	config := flags.StringRequired("c", "configuration", "")
	d := flags.String("d", "debugmessage", "dickhead", "")
	logging := flags.Bool("l", "logging", "")
	log.Println(config)
	log.Println(d)
	log.Println(logging)
}
