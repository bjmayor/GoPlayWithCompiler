package main

import (
	"GoPlayWithCompiler/source/lession03/lexer"
	"GoPlayWithCompiler/source/lession03/parser"
	"fmt"
)

func main() {
	input := `
int a=10;
`

	l := lexer.New(input)
	p := parser.New(l)
	decl := p.Parse()
	fmt.Printf("AST: %s\n\n", decl)
}
