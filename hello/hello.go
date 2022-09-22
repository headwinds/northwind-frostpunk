package hello

import (
	"fmt"

	"rsc.io/quote"
)

func Hello()  string {
    hello := quote.Glass()
    message := fmt.Sprintf("Hi, %v. Welcome!", hello)
    return message
}