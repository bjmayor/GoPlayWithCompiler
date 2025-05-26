package ast

import (
	"GoPlayWithCompiler/source/lession03/token"
	"fmt"
)

// AST 节点接口
type Node interface {
	TokenLiteral() string // 用于调试
	String() string
}

// 语句节点（如 let, return）
type Statement interface {
	Node
	statementNode()
}

// 表达式节点（如 5, add(1,2)）
type Expression interface {
	Node
	expressionNode()
}

// 简单 AST 节点（基础实现）
type SimpleASTNode struct {
	Token      token.Token // 词法单元
	ChildNodes []Node      // 子节点
}

func (n *SimpleASTNode) TokenLiteral() string {
	return n.Token.Literal
}

func NewSimpleASTNode(tok token.Token, children []Node) *SimpleASTNode {
	return &SimpleASTNode{
		Token:      tok,
		ChildNodes: children,
	}
}

// Let 语句节点
type LetStatement struct {
	Token token.Token // LET 令牌
	Name  *Identifier // 变量名
	Value Expression  // 赋值表达式
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	return fmt.Sprintf("%s %s = %s;", ls.TokenLiteral(), ls.Name.String(), ls.Value.String())
}

// 标识符节点
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return fmt.Sprintf("%s", i.Value)
}

// 变量声明节点（如 int a = 10;）
type VarDeclaration struct {
	TypeToken token.Token // int 关键字
	Name      *Identifier // 变量名
	Value     Expression  // 初始值（可选）
}

func (vd *VarDeclaration) statementNode()       {}
func (vd *VarDeclaration) TokenLiteral() string { return vd.TypeToken.Literal }

func (vd *VarDeclaration) String() string {
	return fmt.Sprintf("%s %s = %s;", vd.TypeToken.Literal, vd.Name.String(), vd.Value.String())
}

// 整数字面量节点
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}
