package main

import (
	"GoPlayWithCompiler/antlrdemo/Hello"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseJSONListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := Hello.NewHello(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	stream.Fill()
	for _, token := range stream.GetAllTokens() {
		fmt.Println(token)
	}

}
