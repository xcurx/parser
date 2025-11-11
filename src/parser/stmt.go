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

func parse_block_stmt(p *parser) ast.Stmt {
    p.expect(lexer.OPEN_CURLY)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		body = append(body, parse_stmt(p))
	}

	p.expect(lexer.CLOSE_CURLY)

	return ast.BlockStmt{
		Body: body,
	}
}

func parse_if_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, default_bp)
	consequent := parse_block_stmt(p)

	var alternate ast.Stmt
	if p.currentTokenKind() == lexer.ELSE {
		p.advance()

		if p.currentTokenKind() == lexer.IF {
			alternate = parse_if_stmt(p)
		} else {
			alternate = parse_block_stmt(p)
		}
	}
    
	return ast.IfStmt{
        Condition: condition,
		Consequent: consequent,
		Alternate: alternate,
	}
}

func parse_return_stmt(p *parser) ast.Stmt {
	p.advance()
    stmt := parse_stmt(p)
	// p.expect(lexer.SEMI_COLON)
    _, isExpr := stmt.(ast.ExprStmt)
	_, isFuncLiteral := stmt.(ast.FuncLiteral)
	
	if !isExpr && !isFuncLiteral {
		panic("Invalid expr")
	}

    return ast.ReturnStmt{
		Stmt: stmt,
	}
}

func parse_fn_stmt(p* parser) ast.Stmt {
    p.advance()
	var name string
	if p.currentTokenKind() == lexer.IDENTIFIER {
        name = p.advance().Value
	}
	p.expect(lexer.OPEN_PAREN)
	parameters := make([]ast.Parameter, 0)  

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
        if (len(parameters) != 0) {
			p.expect(lexer.COMMA)
		}

        param_name := p.expect(lexer.IDENTIFIER)
		p.expect(lexer.COLON)
		param_type := parse_type(p, default_bp)
		parameters = append(parameters, ast.Parameter{Name: param_name.Value, Type: param_type})
	}

	p.expect(lexer.CLOSE_PAREN)

	return_type := parse_type(p, default_bp)

	body := parse_block_stmt(p)

	if (name == "") {
		return ast.FuncLiteral{
            Parameter: parameters,
			Return: return_type,
        	Body: body.(ast.BlockStmt).Body,
		}
	}

	return ast.FuncDeclStmt{
        Name: name,
		Parameter: parameters,
		Return: return_type,
        Body: body.(ast.BlockStmt).Body,
	}
}