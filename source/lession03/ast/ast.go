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

// 实现 Node 接口的 String() 方法，递归打印 AST
func (n *SimpleASTNode) String() string {
	result := n.TokenLiteral()
	if len(n.ChildNodes) > 0 {
		result += " ["
		for i, child := range n.ChildNodes {
			if i > 0 {
				result += ", "
			}
			result += child.String()
		}
		result += "]"
	}
	return result
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

// 表达式语句节点
type ExpressionStatement struct {
	Expr Expression
}

func (es *ExpressionStatement) statementNode()      {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Expr.TokenLiteral() }
func (es *ExpressionStatement) String() string       { return es.Expr.String() }

// 二元表达式节点（如 1 + 2, 3 * 4）
type BinaryExpression struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Operator.Literal }
func (be *BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", be.Left.String(), be.Operator.Literal, be.Right.String())
}
