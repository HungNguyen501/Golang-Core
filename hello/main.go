package main

import (
	"fmt"
	"log"

	"golang-core/hello/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	names := []string{
		"Alice",
		"Bob",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
