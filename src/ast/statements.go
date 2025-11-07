package ast

type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExprStmt struct {
	Expression Expr
}

func (n ExprStmt) stmt() {}

type VarDeclStmt struct {
	VariableName  string
	IsConst       bool
	AssignedValue Expr
	ExplicitType  Type
}

func (n VarDeclStmt) stmt() {}
