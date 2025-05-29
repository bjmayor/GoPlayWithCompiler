package main

import (
	"GoPlayWithCompiler/source/simple_calculator/lexer"
	"fmt"
)

func main() {
	lexerInstance := lexer.NewSimpleLexer()

	scripts := []string{
		"int age = 45;",
		"inta age = 45;",
		"in age = 45;",
		"age >= 45;",
		"age > 45;",
	}

	for _, script := range scripts {
		fmt.Printf("parse : %s\n", script)
		tokenReader := lexerInstance.Tokenize(script)
		lexer.Dump(tokenReader)
		fmt.Println()
	}
}
