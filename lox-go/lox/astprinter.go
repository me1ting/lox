package lox

import (
	"bytes"
	"fmt"
)

type AstPrinter struct {
}

func (a *AstPrinter) Print(e Expr) string {
	return e.Accept(a).(string)
}

func (a *AstPrinter) visitBinary(expr *Binary) any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) visitGrouping(expr *Grouping) any {
	return a.parenthesize("group", expr.Expression)
}

func (a *AstPrinter) visitLiteral(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprint(expr.Value)
}

func (a *AstPrinter) visitUnary(expr *Unary) any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var buffer bytes.Buffer
	buffer.WriteByte('(')
	buffer.WriteString(name)
	for _, expr := range exprs {
		buffer.WriteByte(' ')
		buffer.WriteString(expr.Accept(a).(string))
	}
	buffer.WriteByte(')')

	return buffer.String()
}

var _ Visitor = &AstPrinter{}
