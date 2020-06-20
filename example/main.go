package main

import (
	"log"

	"github.com/BlackwonderTF/go-flags"
)

func main() {
	flags.ParseFlags()
	log.Println(*flags.ReadFlag("c").String())
}
