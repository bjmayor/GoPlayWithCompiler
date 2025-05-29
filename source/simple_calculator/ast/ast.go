package ast

// ASTNodeType 表示 AST 节点的类型。
type ASTNodeType string

const (
	Program        ASTNodeType = "Program"        // 程序入口，根节点
	IntDeclaration ASTNodeType = "IntDeclaration" // 整型变量声明
	ExpressionStmt ASTNodeType = "ExpressionStmt" // 表达式语句
	AssignmentStmt ASTNodeType = "AssignmentStmt" // 赋值语句
	Primary        ASTNodeType = "Primary"        // 基础表达式
	Multiplicative ASTNodeType = "Multiplicative" // 乘法表达式
	Additive       ASTNodeType = "Additive"       // 加法表达式
	Identifier     ASTNodeType = "Identifier"     // 标识符
	IntLiteral     ASTNodeType = "IntLiteral"     // 整型字面量
)

// ASTNode 表示抽象语法树的节点接口。
type ASTNode interface {
	Type() ASTNodeType
	Text() string
	Children() []ASTNode
	Parent() ASTNode
	AddChild(child ASTNode)
}

// SimpleASTNode 是 ASTNode 的基础实现。
type SimpleASTNode struct {
	nodeType ASTNodeType
	text     string
	children []ASTNode
	parent   ASTNode
}

func NewSimpleASTNode(nodeType ASTNodeType, text string, parent ASTNode) *SimpleASTNode {
	return &SimpleASTNode{
		nodeType: nodeType,
		text:     text,
		children: []ASTNode{},
		parent:   parent,
	}
}

func (n *SimpleASTNode) Type() ASTNodeType {
	return n.nodeType
}

func (n *SimpleASTNode) Text() string {
	return n.text
}

func (n *SimpleASTNode) Children() []ASTNode {
	return n.children
}

func (n *SimpleASTNode) AddChild(child ASTNode) {
	n.children = append(n.children, child)
}

func (n *SimpleASTNode) SetParent(parent ASTNode) {
	n.parent = parent
}

func (n *SimpleASTNode) Parent() ASTNode {
	return n.parent
}
