package craft

import (
	"fmt"
	"testing"
)

func TestSimpleLexer(t *testing.T)  {
	var lexer  SimpleLexer
	var script string
	var reader TokenReader
	t.Run("45", func(t *testing.T) {
		lexer = SimpleLexer{}
		script = `45`
		fmt.Printf("*************\n parse : %s\n", script)
		reader = lexer.tokenize(script)
		lexer.dump(reader)
	})


	lexer = SimpleLexer{}
	script = `age >= 45;`
	fmt.Printf("*************\n parse : %s\n", script)
	reader = lexer.tokenize(script)
	lexer.dump(reader)

	lexer = SimpleLexer{}
	script = `age > 45;`
	fmt.Printf("*************\n parse : %s\n", script)
	reader = lexer.tokenize(script)
	lexer.dump(reader)


	lexer = SimpleLexer{}
	script = `inta age = 45;`
	fmt.Printf("*************\n parse : %s\n", script)
	reader = lexer.tokenize(script)
	lexer.dump(reader)

	lexer = SimpleLexer{}
	script = `2+3*5;`
	fmt.Printf("*************\n parse : %s\n", script)
	reader = lexer.tokenize(script)
	lexer.dump(reader)
}
