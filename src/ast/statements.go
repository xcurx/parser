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

type IfStmt struct {
	Condition  Expr
	Consequent Stmt
	Alternate  Stmt
}

func (n IfStmt) stmt() {}

type ReturnStmt struct {
	Stmt Stmt
}

func (n ReturnStmt) stmt() {}
type Parameter struct {
	Name string
	Type Type
}

type FuncDeclStmt struct {
	Name      string
	Parameter []Parameter
	Body      []Stmt
	Return    Type
}

func (n FuncDeclStmt) stmt() {}

type FuncLiteral struct {
	Parameter []Parameter
	Body      []Stmt
	Return    Type
}

func (n FuncLiteral) stmt() {}
