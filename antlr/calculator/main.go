package main

import (
	"fmt"

	"calculator/calc"
	"calculator/parsing"

	"github.com/antlr4-go/antlr/v4"
)

func parseTree(prog string) antlr.ParseTree {
	input := antlr.NewInputStream(prog)
	lexer := parsing.NewCalculatorLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parsing.NewCalculatorParser(stream)

	return parser.Prog()
}

func main() {
	prog := `(1 + 2) * 3 - (5 % 3) * 2`
	tree := parseTree(prog)

	// listener := calc.NewCalculatorListener()
	// antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	visitor := calc.NewCalculatorVisitor()
	result := visitor.Visit(tree)
	fmt.Println(result)
}
