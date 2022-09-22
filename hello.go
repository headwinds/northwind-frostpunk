package greetings

import (
	"fmt"

	"rsc.io/quote"
)

func Hello()  string {
    hello := quote.Hello()
    message := fmt.Sprintf("Hi, %v. Welcome!", hello)
    return message
}