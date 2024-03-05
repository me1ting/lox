package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: lox-go [script]")
		os.Exit(64)
	}

	lox := NewLox()
	if len(args) == 1 {
		lox.RunFile(args[0])
	} else {
		lox.Prompt()
	}
}

type Scanner struct{}

func NewScanner(path string) Scanner {
	return Scanner{}
}

func (s *Scanner) ScanTokens() []Token {
	return nil
}

type Token struct{}
