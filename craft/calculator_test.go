package craft

import "testing"

func TestSimpleCalculator_Parse(t *testing.T) {
	var lexer  SimpleLexer
	var script string
	var reader TokenReader


	t.Run("primary", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `45`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.primary(reader)
		DumpAST(*node, "  ")
	})

	t.Run("*", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `32 * 45`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.multiplicative(reader)
		DumpAST(*node, "  ")
	})

	t.Run("+", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `32 + 45`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.addtive(reader)
		DumpAST(*node, "  ")
	})

	t.Run("32+2*5", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `32 + 2 * 5`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.addtive(reader)
		DumpAST(*node, "  ")
	})
	// 结合性有问题。 3+4结合了
	t.Run("addtive1", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `2+3+4`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.addtive1(reader)
		DumpAST(*node, "  ")
	})

	t.Run("addtive", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `2+3+4+5`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.addtive(reader)
		DumpAST(*node, "  ")
	})

	t.Run("evaluate", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `2+3*4`
		result :=  cal.Evaluate(script)
		t.Logf("evaluate %s, result:%d",script, result)
	})

	t.Run("int declare", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `int a;`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.intDeclare(reader)
		DumpAST(*node, "  ")
	})


	t.Run("int declare", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `int a = 35;`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.intDeclare(reader)
		DumpAST(*node, "  ")
	})

	t.Run("int declare", func(t *testing.T) {
		cal := SimpleCalculator{}
		script = `int a = 35 + 47;`
		lexer = SimpleLexer{}
		reader = lexer.tokenize(script)
		node := cal.intDeclare(reader)
		DumpAST(*node, "  ")
	})
}
