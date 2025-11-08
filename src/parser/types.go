package parser

import (
	"fmt"

	"github.com/xcurx/parser/src/ast"
	"github.com/xcurx/parser/src/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}
var type_bp_lu = type_bp_lookup{}

func type_nud(kind lexer.TokenKind, type_nud_fn type_nud_handler) {
	type_nud_lu[kind] = type_nud_fn
}

func type_led(kind lexer.TokenKind, bp binding_power, type_led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = type_led_fn
}

func createTokenTypeLookups() {
    type_nud(lexer.IDENTIFIER, parse_symbol_type)
	type_nud(lexer.OPEN_BRACKET, parse_array_type)
	type_nud(lexer.FN, parse_fn_type)
}

func parse_type(p *parser, bp binding_power) ast.Type {
	//first parse the nud
	tokenKind := p.currentTokenKind()
	type_nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("No type nud function for token kind %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := type_nud_fn(p)

	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		type_led_fn, exists := type_led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("No type led function for token kind %s\n", lexer.TokenKindString(tokenKind)))
		}
        
		left = type_led_fn(p, left, type_bp_lu[tokenKind])
	}

	return left
}


func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CLOSE_BRACKET)
	underlyingType := parse_type(p, default_bp)
	return ast.ArrayType{
		Underlying: underlyingType,
	}
}

func parse_fn_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.OPEN_PAREN)
    args := make([]ast.Type, 0);

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
        if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		}

		args = append(args, parse_type(p, default_bp))
	}

	p.expect(lexer.CLOSE_PAREN)

	returnType := parse_type(p, default_bp)

	parse_block_stmt(p)

    return ast.FuncType{
		Args: args,
		ReturnType: returnType,
	}
} 