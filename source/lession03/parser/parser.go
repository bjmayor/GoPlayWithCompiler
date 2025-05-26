package parser

import (
	"GoPlayWithCompiler/source/lession03/ast"
	"GoPlayWithCompiler/source/lession03/lexer"
	"GoPlayWithCompiler/source/lession03/token"
	"strconv"
)

type Parser struct {
	l       *lexer.Lexer
	errors  []string
	curTok  token.Token
	peekTok token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 读取两个 Token 初始化 curTok 和 peekTok
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.l.NextToken()
}

func (p *Parser) Parse() *ast.SimpleASTNode {
	rootNode := &ast.SimpleASTNode{}
	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			rootNode.ChildNodes = append(rootNode.ChildNodes, stmt)
		}
		p.nextToken()
	}
	return rootNode
}

// 解析语句（示例：只处理 let 语句）
func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.INT_TYPE:
		return p.parseVarDeclaration()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curTok}
	// 这里简化为直接消费 Token，实际需要完整解析 let x = 5; 的结构
	p.nextToken() // 跳过 let
	p.nextToken() // 跳过标识符
	p.nextToken() // 跳过 =
	p.nextToken() // 跳过值
	p.nextToken() // 跳过 ;
	return stmt
}

// 解析 int a = 10; 这种声明
func (p *Parser) parseVarDeclaration() *ast.VarDeclaration {
	decl := &ast.VarDeclaration{TypeToken: p.curTok}

	// 消耗掉 int
	p.nextToken()

	// 检查变量名
	if p.curTok.Type != token.IDENT {
		panic("expected identifier after int")
	}
	decl.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	// 消耗掉变量名
	p.nextToken()

	// 处理可选的初始化部分
	if p.curTok.Type == token.ASSIGN {
		p.nextToken() // 消耗掉 =

		// 解析右侧表达式（简化为只解析整数）
		if p.curTok.Type != token.INT {
			panic("expected integer value")
		}
		val, _ := strconv.ParseInt(p.curTok.Literal, 10, 64)
		decl.Value = &ast.IntegerLiteral{Token: p.curTok, Value: val}

		p.nextToken() // 消耗掉值
	}

	// 检查分号
	if p.curTok.Type != token.SEMICOLON {
		panic("expected semicolon")
	}

	return decl
}
