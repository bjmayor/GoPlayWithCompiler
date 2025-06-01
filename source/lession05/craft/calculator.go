package craft

import (
	"fmt"
	"strconv"
)

type ASTNodeType string
const (
	ASTNodeType_Program  = ASTNodeType("program")
	ASTNodeType_IntLiteral = ASTNodeType("IntLiteral")
	ASTNodeType_IntDeclaration = ASTNodeType("IntDeclaration")
	ASTNodeType_AddtiveExp= ASTNodeType("AddtiveExp")
	ASTNodeType_Multiplicative= ASTNodeType("Multiplicative")
	ASTNodeType_Assignment = ASTNodeType("Assignment")
	ASTNodeType_Identifier= ASTNodeType("Identifier")
)


/**
 * 实现一个计算器，但计算的结合性是有问题的。因为它使用了下面的语法规则：
 *
 * additive -> multiplicative | multiplicative + additive
 * multiplicative -> primary | primary * multiplicative    //感谢@Void_seT，原来写成+号了，写错了。
 *
 * 递归项在右边，会自然的对应右结合。我们真正需要的是左结合。
 */
type SimpleCalculator struct {

}

/*
* 对某个AST节点求值，并打印求值过程。
 */
func (cal *SimpleCalculator)evaluate(node ASTNoder, indent string) int  {
	result := 0
	fmt.Printf("%sCalculating:%s\n", indent, node.GetType())
	switch node.GetType() {
	case ASTNodeType_Program:
		for _, n := range node.GetChildren() {
			result = cal.evaluate(n, indent)
		}
	case ASTNodeType_AddtiveExp:
		child1 := node.GetChildren()[0]
		value1 := cal.evaluate(child1, indent + "\t");
		child2 := node.GetChildren()[1];
		value2 := cal.evaluate(child2, indent + "\t");
		if (node.GetText() == "+") {
			result = value1 + value2;
		} else {
			result = value1 - value2;
		}
	case ASTNodeType_IntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	case ASTNodeType_Multiplicative:
		child1 := node.GetChildren()[0]
		value1 := cal.evaluate(child1, indent + "\t");
		child2 := node.GetChildren()[1];
		value2 := cal.evaluate(child2, indent + "\t");
		if (node.GetText() == "*") {
			result = value1 * value2;
		} else {
			result = value1 / value2;
		}
	}
	return result
}

/**
     * 语法解析：根节点
 */
func (cal *SimpleCalculator)prog(tokens TokenReader)  *ASTNoder {
	noder := NewASTNoder(ASTNodeType_Program, "Calculator")
	child := cal.addtive(tokens)
	if child != nil {
		noder.AddChild(*child)
	}
	return &noder
}

/**
 * 整型变量声明语句
 */
func (cal *SimpleCalculator)intDeclare(reader TokenReader) *ASTNoder  {
	var node ASTNoder
	token :=  reader.Peek()
	if token!=nil && token.Type == TokenType_Int {
		reader.Read()
		if reader.Peek().Type == TokenType_Id {
			token := reader.Read()
			node = NewASTNoder(ASTNodeType_IntDeclaration, token.Text)
			token = reader.Peek()
			if token != nil && token.Type == TokenType_Assignment {
				reader.Read()
				child := cal.addtive(reader)
				if child == nil {
					panic("invalide variable initialization, expecting an expression")
				} else {
					node.AddChild(*child)
				}
			}
		} else {
			panic("variable name expected")
		}
		if node != nil {
			token := reader.Peek()
			if token!=nil && token.Type == TokenType_SemiColon {
				reader.Read()
			} else {
				panic("invalid statement, expecting semicolon")
			}
		}
	}
	return &node
}

/**
 * 语法解析：加法表达式
add -> mul
add -> mul + add
 */
func (cal *SimpleCalculator)addtive1(reader TokenReader) *ASTNoder  {
	child1 := cal.multiplicative(reader)
	var node ASTNoder
	token := reader.Peek()
	if child1!=nil && token!=nil {
		if token.Type == TokenType_Plus {
			token := reader.Read()
			child2 := cal.addtive(reader)
			if child2!=nil {
				node = NewASTNoder(ASTNodeType_AddtiveExp,token.Text)
				node.AddChild(*child1)
				node.AddChild(*child2)
			} else {
				panic("invalid additive expression, expecting the right part.")
			}
		}
	}
	if node != nil {
		return &node
	}
	return child1
}

/**
 * 语法解析：加法表达式
add -> mul ( + mul ) *
*/
func (cal *SimpleCalculator)addtive(reader TokenReader) *ASTNoder  {
	child1 := cal.multiplicative(reader)
	var node ASTNoder
	if child1!=nil {
		for {
			token := reader.Peek()
			if token != nil && token.Type == TokenType_Plus {
				token := reader.Read()
				child2 := cal.multiplicative(reader)
				if child2!=nil {
					node = NewASTNoder(ASTNodeType_AddtiveExp,token.Text)
					node.AddChild(*child1)
					node.AddChild(*child2)
					*child1 = node
				} else {
					panic("invalid additive expression, expecting the right part.")
				}
			} else {
				break
			}
		}

	}
	return child1
}


/**
 * 语法解析：基础表达式
 */
func (cal *SimpleCalculator)primary(reader TokenReader) *ASTNoder  {
	var node ASTNoder
	token := reader.Peek()
	if token!= nil {
		switch token.Type {
		case TokenType_IntLiteral:
			reader.Read()
			node = NewASTNoder(ASTNodeType_IntLiteral, token.Text)
		case TokenType_Id:
			token := reader.Read()
			node = NewASTNoder(ASTNodeType_IntDeclaration, token.Text)
		case TokenType_Left_Paren:
			reader.Read()
			node := cal.addtive(reader)
			if node != nil {
				token := reader.Peek()
				if token!= nil && token.Type == TokenType_Right_Paren {
					reader.Read()
				} else {
					panic("expecting right parenthesis")
				}
			} else {
				panic("expecting an additive expression inside parenthesis")
			}
		}
	}
	return &node
}

/**
 * 语法解析：乘法表达式
 */
func (cal *SimpleCalculator)multiplicative(reader TokenReader) *ASTNoder  {
	child1 := cal.primary(reader);
	var node ASTNoder
	token := reader.Peek()
	if child1!= nil && token != nil {
		if token.Type == TokenType_Star {
			token := reader.Read()
			child2 := cal.primary(reader)
			if child2 != nil {
				node = NewASTNoder(ASTNodeType_Multiplicative, token.Text)
				node.AddChild(*child1)
				node.AddChild(*child2)
			} else {
				panic("invalid multiplicative expression, expecting the right part.")
			}
		}
	}
	if node != nil {
		return &node
	}
	return child1
}

func (cal *SimpleCalculator)Evaluate(script string) int {
	node := cal.Parse(script)
	DumpAST(*node,"  ")
	return cal.evaluate(*node, "  ")
}

func (cal *SimpleCalculator)Parse(code string)*ASTNoder {
	lexer := SimpleLexer{}
	tokens := lexer.tokenize(code)
	return cal.prog(tokens)
}


type ASTNoder interface {
	AddChild(child ASTNoder)
	GetText() string
	GetType() ASTNodeType
	GetChildren() []ASTNoder
	GetParent() ASTNoder
}
type SimpleASTNode struct {
	nodeType ASTNodeType
	text string
	parent ASTNoder
	children []ASTNoder
}

func (s *SimpleASTNode) AddChild(child ASTNoder) {
	s.children = append(s.children, child)
}

func (s *SimpleASTNode) GetText() string {
	return s.text
}

func (s *SimpleASTNode) GetType() ASTNodeType {
	return s.nodeType
}

func (s *SimpleASTNode) GetChildren() []ASTNoder {
	return s.children
}

func (s *SimpleASTNode) GetParent() ASTNoder {
	return s.parent
}

func DumpAST(node ASTNoder, indent string) {
	fmt.Printf("%s%s %s\n", indent, node.GetType(), node.GetText())
	for _, _node := range node.GetChildren(){
		_node := _node
		DumpAST(_node, "\t"+indent)
	}
}

func NewASTNoder(nodeType ASTNodeType, text string)  ASTNoder {
	return &SimpleASTNode{nodeType:nodeType, text:text}
}