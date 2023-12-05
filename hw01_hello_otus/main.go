package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	initialString := "Hello, OTUS!"
	reversedString := reverse.String(initialString)

	fmt.Println(reversedString)
}
