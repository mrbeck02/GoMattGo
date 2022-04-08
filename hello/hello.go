package main

import (
	"fmt"
	"log"

	"github.com/mrbeck02/GoMattGo/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("loggy: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ladies")
	fmt.Println(message)

	fmt.Println(quote.Go())
}
