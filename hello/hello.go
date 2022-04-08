package main

import (
	"fmt"

	"github.com/mrbeck02/GoMattGo/greetings"
	"rsc.io/quote"
)

func main() {
	message := greetings.Hello("Ladies")
	fmt.Println(message)

	fmt.Println(quote.Go())
}
