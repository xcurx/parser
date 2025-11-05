package parser

import (
	"github.com/xcurx/parser/src/ast"
	"github.com/xcurx/parser/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens,
	}
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
    body := make([]ast.Stmt, 0)
    p := createParser(tokens)
    
	for p.hasTokens() {
        body = append(body, parse_stmt(p))	
	}

	return ast.BlockStmt{
		Body: body,
	}
}

//helpers
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {
    token := p.currentToken()
	p.pos++
	return token
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}
