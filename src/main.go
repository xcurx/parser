package main

import (
	"os"

	"github.com/sanity-io/litter"
	"github.com/xcurx/parser/src/lexer"
	"github.com/xcurx/parser/src/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/04.lang")
	source := string(bytes)
	
	tokens := lexer.Tokenize(source)

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}