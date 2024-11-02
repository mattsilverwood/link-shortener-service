package main

import (
	"log"

	"github.com/mattsilverwood/link-shortener-service/cmd"
)

func main() {
	log.Fatal(cmd.StartService())
}
