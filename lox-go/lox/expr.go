package lox

type Expr interface {
	Accept(Visitor) any
}

type Visitor interface {
	visitBinary(*Binary) any
	visitGrouping(*Grouping) any
	visitLiteral(*Literal) any
	visitUnary(*Unary) any
}

type Binary struct {
	Left     Expr
	Operator *Token
	Right    Expr
}

func (b *Binary) Accept(v Visitor) any {
	return v.visitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(v Visitor) any {
	return v.visitGrouping(g)
}

type Literal struct {
	Value any
}

func (li *Literal) Accept(v Visitor) any {
	return v.visitLiteral(li)
}

type Unary struct {
	Operator *Token
	Right    Expr
}

func (u *Unary) Accept(v Visitor) any {
	return v.visitUnary(u)
}

var _ Expr = &Binary{}
var _ Expr = &Grouping{}
var _ Expr = &Literal{}
var _ Expr = &Unary{}
