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
	case token.INT_TYPE:
		return p.parseVarDeclaration()
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

// 解析乘法表达式（* 和 /），递归处理左结合
func (p *Parser) multiplicative() *ast.SimpleASTNode {
	child1 := p.primary()
	node := child1

	for {
		if node != nil && (p.curTok.Type == token.ASTERISK || p.curTok.Type == token.SLASH) {
			opTok := p.curTok
			p.nextToken()
			child2 := p.primary()
			if child2 != nil {
				newNode := ast.NewSimpleASTNode(opTok, []ast.Node{})
				newNode.ChildNodes = append(newNode.ChildNodes, node)
				newNode.ChildNodes = append(newNode.ChildNodes, child2)
				node = newNode
			} else {
				return nil
			}
		} else {
			break
		}
	}
	return node
}

// 解析加法表达式（+），递归处理左结合
func (p *Parser) additive() *ast.SimpleASTNode {
	child1 := p.multiplicative()
	node := child1

	for {
		if node != nil && p.curTok.Type == token.PLUS {
			opTok := p.curTok
			p.nextToken()
			child2 := p.additive()
			if child2 != nil {
				newNode := ast.NewSimpleASTNode(opTok, []ast.Node{})
				newNode.ChildNodes = append(newNode.ChildNodes, node)
				newNode.ChildNodes = append(newNode.ChildNodes, child2)
				node = newNode
			} else {
				return nil
			}
		} else {
			break
		}
	}
	return node
}

// 解析基础表达式（整数字面量、括号等）
func (p *Parser) primary() *ast.SimpleASTNode {
	if p.curTok.Type == token.INT {
		tok := p.curTok
		p.nextToken()
		node := ast.NewSimpleASTNode(tok, []ast.Node{})
		return node
	} else if p.curTok.Type == token.LPAREN {
		p.nextToken() // 消耗 (
		node := p.additive()
		if p.curTok.Type == token.RPAREN {
			p.nextToken() // 消耗 )
			return node
		} else {
			return nil // 缺少右括号
		}
	}
	return nil
}
