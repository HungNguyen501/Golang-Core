package main

import (
	"fmt"
	"log"

	"golang-core/hello-go/src"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	names := []string{
		"Alice",
		"Bob",
	}
	messages, err := src.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
