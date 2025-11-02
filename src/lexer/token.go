package lexer

import (
	"fmt"
	"slices"
)

type Token struct {
	Kind TokenKind
	Value string
}

func (token Token) isOneOfMany(kinds ...TokenKind) bool {
	return slices.Contains(kinds, token.Kind)
}

func (token Token) Debug() {
	if (token.isOneOfMany(IDENTIFIER, NUMBER, STRING)) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}