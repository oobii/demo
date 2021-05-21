package main

import (
	"fmt"
	"log"

	"example.com/demo/mascot"
	"example.com/greetings2"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings2: ")
	log.SetFlags(0)
	
	fmt.Println(mascot.BestMascot())
	fmt.Println(mascot.Greeting("Hello"))
	fmt.Println(quote.Go())

	// fmt.Println(greetings2.Hello("Joe"))

	names := []string {"Joe", "Tim", "Sara"}
	
	messages, err := greetings2.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
