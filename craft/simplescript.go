package craft

import (
	"fmt"
	"strconv"
)

type SimpleScript struct {
	variables map[string]int
	verbose bool
}

func NewSimpleScript(verbose bool) *SimpleScript {
	return &SimpleScript{
		variables:make(map[string]int),
		verbose:verbose,
	}
}

func (s *SimpleScript) Evaluate(node ASTNoder, indent string) int {
	result := 0
	if s.verbose {
		fmt.Printf("%sCalculating:%s\n", indent, node.GetType())
	}
	switch node.GetType() {
	case ASTNodeType_Program:
		for _, n := range node.GetChildren() {
			result = s.Evaluate(n, indent)
		}
	case ASTNodeType_AddtiveExp:
		child1 := node.GetChildren()[0]
		value1 := s.Evaluate(child1, indent + "\t");
		child2 := node.GetChildren()[1];
		value2 := s.Evaluate(child2, indent + "\t");
		if (node.GetText() == "+") {
			result = value1 + value2;
		} else {
			result = value1 - value2;
		}
	case ASTNodeType_IntLiteral:
		result, _ = strconv.Atoi(node.GetText())
	case ASTNodeType_Multiplicative:
		child1 := node.GetChildren()[0]
		value1 := s.Evaluate(child1, indent + "\t");
		child2 := node.GetChildren()[1];
		value2 := s.Evaluate(child2, indent + "\t");
		if (node.GetText() == "*") {
			result = value1 * value2;
		} else {
			result = value1 / value2;
		}
	case ASTNodeType_Identifier:
		varName := node.GetText()
		if v, ok := s.variables[varName]; !ok {
			panic("unknown variable: " + varName)
		} else {
			result = v
		}
	case ASTNodeType_Assignment:
		varName := node.GetText()
		if _, ok := s.variables[varName]; !ok {
			panic("unknown variable: " + varName)
		}
	case ASTNodeType_IntDeclaration:
		varName := node.GetText()
		varValue := 0
		if len(node.GetChildren()) > 0 {
			child := node.GetChildren()[0]
			result := s.Evaluate(child, indent+"\t")
			varValue = result
		}
		s.variables[varName] = varValue
	}

	if s.verbose {
		fmt.Println(indent, "result:", result)
	} else if indent == "" {
		switch node.GetType() {
		case ASTNodeType_IntDeclaration:
			fallthrough
		case ASTNodeType_Assignment:
			fmt.Println(node.GetText(), "result:", result)
		case ASTNodeType_Program:
		default:
			fmt.Println("result:", result)
		}
	}
	return result
}
