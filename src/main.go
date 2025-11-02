package main

import (
	// "fmt"
	"os"
	"github.com/xcurx/parser/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	tokens := lexer.Tokenize(string(source))

	for _, token := range tokens {
		token.Debug()
	}
}