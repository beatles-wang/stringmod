package main

import (
	"fmt"

	"github.com/beatles-wang/stringmod/quotes"
	"github.com/beatles-wang/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("123456789")
	fmt.Println(o, e)

	fmt.Println(quotes.GetEmoji("turtle"))
}
