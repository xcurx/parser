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
    isConst := p.advance().Kind == lexer.CONST
    varName := p.expectError(lexer.IDENTIFIER, "Inside variable decleration expected to find varible name").Value
    p.expect(lexer.ASSIGNMENT)
    assignedValue := parse_expr(p, assignment)
    p.expect(lexer.SEMI_COLON)

    return ast.VarDeclStmt{
        VariableName: varName,
        IsConst: isConst,
        AssignedValue: assignedValue,
    }
}