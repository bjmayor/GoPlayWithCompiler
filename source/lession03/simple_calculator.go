package main

import (
	"GoPlayWithCompiler/source/lession03/lexer"
	"GoPlayWithCompiler/source/lession03/parser"
	"fmt"
)

func main() {
	input := `
int a=10;
2+3*5;
`

	l := lexer.New(input) // 解析成token流 l.newToken()
	p := parser.New(l)
	decl := p.Parse() // 解析成ast树
	fmt.Printf("AST: %s\n\n", decl)
}
