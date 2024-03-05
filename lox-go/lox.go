package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Lox struct {
	hadError bool
}

func NewLox() Lox {
	return Lox{}
}

// 执行脚本
func (x *Lox) RunFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	x.Run(string(data))
	if x.hadError {
		os.Exit(65)
	}
}

// 交互式解释器
func (x *Lox) Prompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			x.Run(line)
			x.hadError = false
		}
	}
}

func (x *Lox) Run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()

	fmt.Println(tokens)
}

func (x *Lox) Error(line int, msg string) {
	x.Report(line, "", msg)
}

func (x *Lox) Report(line int, where, msg string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, msg)
	x.hadError = true
}
