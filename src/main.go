package main

import (
	"fmt"
	"os"
	"parser/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	fmt.Printf("Code: %s\n", source)
}
