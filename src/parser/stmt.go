package parser

import (
	"github.com/xcurx/parser/src/ast"
	"github.com/xcurx/parser/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	}

	expr := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExprStmt{
		Expression: expr,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
    var explicitType ast.Type
    var assignedValue ast.Expr
	isConst := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable decleration expected to find varible name").Value

    //check for explicit type
    if p.currentTokenKind() == lexer.COLON {
        p.advance()
        explicitType = parse_type(p, default_bp)
    }

    if p.currentTokenKind() != lexer.SEMI_COLON {
        p.expect(lexer.ASSIGNMENT)
        assignedValue = parse_expr(p, assignment)
    } else if explicitType == nil {
        panic("Variable declaration without assigned value must have explicit type")
    }

    p.expect(lexer.SEMI_COLON)

    if isConst && assignedValue == nil {
        panic("Const variable declaration must have assigned value")
    }

	return ast.VarDeclStmt{
		VariableName:  varName,
		IsConst:       isConst,
		AssignedValue: assignedValue,
		ExplicitType:  explicitType,
	}
}
