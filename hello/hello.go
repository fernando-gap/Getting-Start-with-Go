package main

import (
	"fmt"
	"log"
	"xica/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	names := []string {
		"Fernando",
		"Augusto",
		"Jerry",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
