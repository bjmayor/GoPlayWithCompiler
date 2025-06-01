package craft

import (
	"bytes"
	"fmt"
)
type DfaState int
const (
	DfaState_Initial DfaState=  iota
	DfaState_Id
	DfaState_Int1
	DfaState_Int2
	DfaState_Int3
	DfaState_Assignment
	DfaState_SemiColon
	DfaState_Left_Paren
	DfaState_Right_Paren
	DfaState_GT
	DfaState_GE
	DfaState_Plus
	DfaState_Minus
	DfaState_Star
	DfaState_Slash
	DfaState_IntLiteral
)

type TokenType string
const (
	TokenType_Initial =  TokenType("Initial")
	TokenType_Id = TokenType("Identifier")
	TokenType_GT = TokenType("GT")
	TokenType_GE = TokenType("GE")
	TokenType_IntLiteral = TokenType("IntLiteral")
	TokenType_Int= TokenType("Int")
	TokenType_Assignment= TokenType("Assignment")
	TokenType_SemiColon= TokenType("SemiColon")
	TokenType_Plus= TokenType("Plus")
	TokenType_Minus = TokenType("Minus")
	TokenType_Star = TokenType("Star")
	TokenType_Slash= TokenType("Slash")
	TokenType_Left_Paren = TokenType("(")
	TokenType_Right_Paren = TokenType(")")
)

type TokenReader interface {
	Read() *Token
	Peek() *Token
	UnRead()
	GetPosition() int
	setPosition(position int)
}

func isAlpha(ch int32) bool  {
	return (ch >= 'a' && ch <= 'z') || (ch>='A' && ch<='Z')
}
func isDigit(ch int32) bool {
	return ch >= '0' && ch <='9'
}

type Token struct {
	Text string
	Type TokenType
}

type SimpleLexer struct {
	tokens []Token
	token Token
	tokenText *bytes.Buffer
}

func (lexer *SimpleLexer)tokenize(script string) TokenReader  {
	lexer.tokenText = new(bytes.Buffer)
	state := DfaState_Initial
	for _, ch := range script {
		//fmt.Printf("%c\n",ch)
		switch state {
		case DfaState_Initial:
			state = lexer.initToken(ch)
		case DfaState_Id:
			if isAlpha(ch) || isDigit(ch) {
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.initToken(ch)
			}
		case DfaState_Int1:
			if ch == 'n' {
				state = DfaState_Int2
				lexer.tokenText.WriteRune(ch)
			} else if isAlpha(ch) || isDigit(ch){
				lexer.tokenText.WriteRune(ch)
				state = DfaState_Id
			} else {
				state = lexer.initToken(ch)
			}
		case DfaState_Int2:
			if ch == 't' {
				state = DfaState_Int3
				lexer.tokenText.WriteRune(ch)
			}else if isAlpha(ch) || isDigit(ch){
				lexer.tokenText.WriteRune(ch)
				state = DfaState_Id
			} else {
				state = lexer.initToken(ch)
			}
		case DfaState_Int3:
			if ch == ' ' {
				lexer.token.Type = TokenType_Int
				state = lexer.initToken(ch)
			} else if isAlpha(ch) || isDigit(ch){
				lexer.tokenText.WriteRune(ch)
				state = DfaState_Id
			} else {
				state = lexer.initToken(ch)
			}

		case DfaState_GT:
			if ch == '=' {
				lexer.token.Type = TokenType_GE
				state = DfaState_GE
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.initToken(ch)
			}
		case DfaState_GE:
			state = lexer.initToken(ch)
		case DfaState_Assignment:
			state = lexer.initToken(ch)
		case DfaState_Plus:
			state = lexer.initToken(ch)
		case DfaState_Minus:
			state = lexer.initToken(ch)
		case DfaState_Star:
			state = lexer.initToken(ch)
		case DfaState_Slash:
			state = lexer.initToken(ch)
		case DfaState_SemiColon:
			state = lexer.initToken(ch)
		case DfaState_Left_Paren:
			state = lexer.initToken(ch)
		case DfaState_Right_Paren:
			state = lexer.initToken(ch)
		case DfaState_IntLiteral:
			if isDigit(ch) {
				lexer.tokenText.WriteRune(ch)
			} else {
				state = lexer.initToken(ch)
			}
		}
	}
	if len(lexer.tokenText.Bytes()) > 0 {
		lexer.initToken(' ')
	}
	return NewTokenReader(lexer.tokens)
}

func (lexer *SimpleLexer)initToken(ch rune) DfaState {
	if len(lexer.tokenText.Bytes()) > 0 {
		lexer.token.Text = lexer.tokenText.String()
		lexer.tokens = append(lexer.tokens, lexer.token)
	}
	lexer.tokenText = new(bytes.Buffer)
	lexer.token = Token{}
	newstate := DfaState_Initial
	switch {
	case isAlpha(ch):
		if ch == 'i' {
			newstate = DfaState_Int1
		} else {
			newstate = DfaState_Id
		}
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Id
	case isDigit(ch):
		newstate = DfaState_IntLiteral
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_IntLiteral
	case ch == '>':
		newstate = DfaState_GT
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_GT
	case ch == '=':
		newstate = DfaState_Assignment
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Assignment
	case ch == '+':
		newstate = DfaState_Plus
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Plus
	case ch == '-':
		newstate = DfaState_Minus
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Minus
	case ch == '*':
		newstate = DfaState_Star
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Star
	case ch == '/':
		newstate = DfaState_Slash
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Slash
	case ch == ';':
		newstate = DfaState_SemiColon
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_SemiColon
	case ch == '(':
		newstate = DfaState_Left_Paren
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Left_Paren
	case ch == ')':
		newstate = DfaState_Right_Paren
		lexer.tokenText.WriteRune(ch)
		lexer.token.Type = TokenType_Right_Paren
	}
	return newstate
}

func (lexer *SimpleLexer)dump(reader TokenReader) {
	fmt.Println("text\ttype")
	var token  *Token
	for  {
		if token = reader.Read();token==nil {
			break
		} else {
			fmt.Printf("%s\t\t%s\n",(*token).Text,(*token).Type)
		}
	}
}



type SimpleTokenReader struct {
	tokens []Token
	position int
}

func (s *SimpleTokenReader) Read() *Token {
	if s.position<len(s.tokens)	{
		p := s.position
		s.position ++
		return &s.tokens[p]
	}
	return nil
}

func (s *SimpleTokenReader) Peek() *Token {
	if s.position<len(s.tokens)	{
		p := s.position
		return &s.tokens[p]
	}
	return nil
}

func (s *SimpleTokenReader) UnRead() {
	if s.position > 0 {
		s.position--
	}
}

func (s *SimpleTokenReader) GetPosition() int {
	return s.position
}

func (s *SimpleTokenReader) setPosition(position int) {
	s.position = position
}

func NewTokenReader(tokens []Token) TokenReader {
	return &SimpleTokenReader{tokens:tokens}
}