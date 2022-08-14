package main

import (
	"github.com/Jiacheng787/create-rc/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
