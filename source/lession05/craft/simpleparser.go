package craft

import (
	"fmt"
	"strconv"
)

/**
 * 一个简单的语法解析器。
 * 能够解析简单的表达式、变量声明和初始化语句、赋值语句。
 * 它支持的语法规则为：
 *
 * programm -> intDeclare | expressionStatement | assignmentStatement
 * intDeclare -> 'int' Id ( = additive) ';'
 * expressionStatement -> addtive ';'
 * addtive -> multiplicative ( (+ | -) multiplicative)*
 * multiplicative -> primary ( (* | /) primary)*
 * primary -> IntLiteral | Id | (additive)
 */
type SimpleParser struct {
}

/*
* 对某个AST节点求值，并打印求值过程。
 */
func (cal *SimpleParser) evaluate(node ASTNoder, indent string) int {
	result := 0
	fmt.Printf("%sCalculating:%s\n", indent, node.GetType())
	switch node.GetType() {
	case ASTNodeType_Program:
		for _, n := range node.GetChildren() {
			result = cal.evaluate(n, indent)
		}
	case ASTNodeType_AddtiveExp:
		child1 := node.GetChildren()[0]
		value1 := cal.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := cal.evaluate(child2, indent+"\t")
		if node.GetText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case ASTNodeType_IntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	case ASTNodeType_Multiplicative:
		child1 := node.GetChildren()[0]
		value1 := cal.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := cal.evaluate(child2, indent+"\t")
		if node.GetText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	}
	return result
}

/**
 * 语法解析：根节点
 */
func (cal *SimpleParser) prog(tokens TokenReader) *ASTNoder {
	noder := NewASTNoder(ASTNodeType_Program, "pwc")
	for {
		token := tokens.Peek()
		if token == nil {
			break
		}
		child := cal.intDeclare(tokens)
		if child == nil {
			child = cal.expressionStatement(tokens)
		}
		if child == nil {
			child = cal.assignmentStatement(tokens)
		}
		if child != nil {
			noder.AddChild(*child)
		} else {
			panic("unknown statement")
		}

	}
	return &noder
}

/**
 * 整型变量声明语句
 */
func (cal *SimpleParser) intDeclare(reader TokenReader) *ASTNoder {
	var node ASTNoder
	token := reader.Peek()
	if token != nil && token.Type == TokenType_Int {
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
			if token != nil && token.Type == TokenType_SemiColon {
				reader.Read()
			} else {
				panic("invalid statement, expecting semicolon")
			}
		}
	}
	if node != nil {
		return &node
	}
	return nil
}

/**
 * 语法解析：加法表达式
add -> mul
add -> mul + add
*/
func (cal *SimpleParser) addtive1(reader TokenReader) *ASTNoder {
	child1 := cal.multiplicative(reader)
	var node ASTNoder
	token := reader.Peek()
	if child1 != nil && token != nil {
		if token.Type == TokenType_Plus {
			token := reader.Read()
			child2 := cal.addtive(reader)
			if child2 != nil {
				node = NewASTNoder(ASTNodeType_AddtiveExp, token.Text)
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
func (cal *SimpleParser) addtive(reader TokenReader) *ASTNoder {
	child1 := cal.multiplicative(reader)
	var node ASTNoder
	if child1 != nil {
		for {
			token := reader.Peek()
			if token != nil && token.Type == TokenType_Plus {
				token := reader.Read()
				child2 := cal.multiplicative(reader)
				if child2 != nil {
					node = NewASTNoder(ASTNodeType_AddtiveExp, token.Text)
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
func (cal *SimpleParser) expressionStatement(reader TokenReader) *ASTNoder {
	pos := reader.GetPosition()
	node := cal.addtive(reader)
	if node != nil {
		token := reader.Peek()
		if token != nil && token.Type == TokenType_SemiColon {
			reader.Read()
		} else {
			node = nil
			reader.setPosition(pos)
		}
	}
	return node
}

func (cal *SimpleParser) assignmentStatement(reader TokenReader) *ASTNoder {
	var node ASTNoder
	token := reader.Peek()
	if token != nil && token.Type == TokenType_Id {
		token := reader.Read()
		node = NewASTNoder(ASTNodeType_Assignment, token.Text)
		token = reader.Peek()
		if token != nil && token.Type == TokenType_Assignment {
			reader.Read()
			child := cal.addtive(reader)
			if child == nil {
				panic("invalide assignment statement, expecting an expression")
			} else {
				node.AddChild(*child)
				token := reader.Peek()
				if token != nil && token.Type == TokenType_SemiColon {
					reader.Read()
				} else {
					panic("invalid statement, expecting semicolon")
				}
			}
		} else {
			reader.UnRead()
			return nil
		}
	}
	return &node
}

/**
 * 语法解析：基础表达式
 */
func (cal *SimpleParser) primary(reader TokenReader) *ASTNoder {
	var node ASTNoder
	token := reader.Peek()
	if token != nil {
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
				if token != nil && token.Type == TokenType_Right_Paren {
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
func (cal *SimpleParser) multiplicative(reader TokenReader) *ASTNoder {
	child1 := cal.primary(reader)
	var node ASTNoder
	token := reader.Peek()
	if child1 != nil && token != nil {
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

func (cal *SimpleParser) Evaluate(script string) int {
	node := cal.Parse(script)
	DumpAST(*node, "  ")
	return cal.evaluate(*node, "  ")
}

func (cal *SimpleParser) Parse(code string) *ASTNoder {
	lexer := SimpleLexer{}
	tokens := lexer.tokenize(code)
	return cal.prog(tokens)
}
