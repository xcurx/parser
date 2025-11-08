package ast

type SymbolType struct {
    Name string
}

func (t SymbolType) _type() {}

type ArrayType struct {
    Underlying Type
}

func (t ArrayType) _type() {}

type FuncType struct {
    Args []Type
    ReturnType Type
}

func (t FuncType) _type() {}