package main

import (
	parser "GoPlayWithCompiler/antlrdemo"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BasePlayScriptListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewPlayScriptLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPlayScriptParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Expression()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
