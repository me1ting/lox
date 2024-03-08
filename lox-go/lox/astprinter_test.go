package lox_test

import (
	. "lox-go/lox"
	"testing"
)

func TestIt(t *testing.T) {
	expr := &Binary{&Unary{
		&Token{MINUS, "-", nil, 1},
		&Literal{123},
	},
		&Token{STAR, "*", nil, 1},
		&Grouping{&Literal{45.67}}}

	if new(AstPrinter).Print(expr) != "(* (- 123) (group 45.67))" {
		t.Fatal("error")
	}
}
