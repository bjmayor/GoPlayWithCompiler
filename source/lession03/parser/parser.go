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

// 新增：表达式语句节点和解析方法
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expr := p.parseAdditiveExpression()
	// 可选：检查分号
	if p.curTok.Type == token.SEMICOLON {
		p.nextToken()
	}
	return &ast.ExpressionStatement{Expr: expr}
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

// 解析语句（支持变量声明和表达式语句）
func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.INT_TYPE:
		return p.parseVarDeclaration()
	case token.INT, token.LPAREN:
		// 以数字或左括号开头的，解析为表达式语句
		return p.parseExpressionStatement()
	default:
		return nil
	}
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
		decl.Value = p.parseAdditiveExpression()
	}

	// 检查分号
	if p.curTok.Type != token.SEMICOLON {
		panic("expected semicolon")
	}

	return decl
}

// ===== 新增：表达式解析相关 =====

// 解析加法表达式
func (p *Parser) parseAdditiveExpression() ast.Expression {
	left := p.parseMultiplicativeExpression()
	for p.curTok.Type == token.PLUS || p.curTok.Type == token.MINUS {
		op := p.curTok
		p.nextToken()
		right := p.parseMultiplicativeExpression()
		left = &ast.BinaryExpression{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// 解析乘法表达式
func (p *Parser) parseMultiplicativeExpression() ast.Expression {
	left := p.parsePrimary()
	for p.curTok.Type == token.ASTERISK || p.curTok.Type == token.SLASH {
		op := p.curTok
		p.nextToken()
		right := p.parsePrimary()
		left = &ast.BinaryExpression{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left
}

// 解析基础表达式（整数字面量、括号等）
func (p *Parser) parsePrimary() ast.Expression {
	if p.curTok.Type == token.INT {
		val, _ := strconv.ParseInt(p.curTok.Literal, 10, 64)
		node := &ast.IntegerLiteral{Token: p.curTok, Value: val}
		p.nextToken()
		return node
	} else if p.curTok.Type == token.LPAREN {
		p.nextToken()
		exp := p.parseAdditiveExpression()
		if p.curTok.Type != token.RPAREN {
			panic("expected )")
		}
		p.nextToken()
		return exp
	}
	panic("invalid primary expression")
}
