package main

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	GT        = ">"
	GE        = ">="
	INT_TYPE  = "INT_TYPE"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	MUL       = "*"
	DIV       = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

type Token struct {
	Type    TokenType
	Literal string
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

type DfaState int

const (
	Initial DfaState = iota
	Id
	Int1
	Int2
	Int3
	GTState
	GEState
	IntLiteral
	Assign
	Plus
	Minus
	Mul
	Div
	SemiColon
	LP
	RP
	LB
	RB
)

var keywords = map[string]TokenType{
	"int": INT_TYPE,
}

type Lexer struct {
	input  []rune
	pos    int
	state  DfaState
	tokens []Token
}

func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
		pos:   0,
		state: Initial,
	}
}

func (l *Lexer) readChar() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	ch := l.input[l.pos]
	l.pos++
	return ch
}

func (l *Lexer) peekChar() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) NextToken() Token {
	var tokenText []rune
	state := Initial
	var ch rune

	for {
		if l.pos >= len(l.input) {
			if len(tokenText) > 0 {
				return l.makeToken(state, string(tokenText))
			}
			return Token{Type: EOF, Literal: ""}
		}
		ch = l.readChar()
		switch state {
		case Initial:
			if isBlank(ch) {
				continue
			} else if isLetter(ch) {
				tokenText = append(tokenText, ch)
				if ch == 'i' {
					state = Int1
				} else {
					state = Id
				}
			} else if isDigit(ch) {
				tokenText = append(tokenText, ch)
				state = IntLiteral
			} else {
				switch ch {
				case '>':
					tokenText = append(tokenText, ch)
					state = GTState
				case '=':
					return Token{Type: ASSIGN, Literal: string(ch)}
				case '+':
					return Token{Type: PLUS, Literal: string(ch)}
				case '-':
					return Token{Type: MINUS, Literal: string(ch)}
				case '*':
					return Token{Type: MUL, Literal: string(ch)}
				case '/':
					return Token{Type: DIV, Literal: string(ch)}
				case ';':
					return Token{Type: SEMICOLON, Literal: string(ch)}
				case '(':
					return Token{Type: LPAREN, Literal: string(ch)}
				case ')':
					return Token{Type: RPAREN, Literal: string(ch)}
				case '{':
					return Token{Type: LBRACE, Literal: string(ch)}
				case '}':
					return Token{Type: RBRACE, Literal: string(ch)}
				default:
					return Token{Type: ILLEGAL, Literal: string(ch)}
				}
			}
		case Id:
			if isLetter(ch) || isDigit(ch) {
				tokenText = append(tokenText, ch)
			} else {
				l.pos-- // unread
				return l.makeToken(Id, string(tokenText))
			}
		case Int1:
			if ch == 'n' {
				tokenText = append(tokenText, ch)
				state = Int2
			} else if isLetter(ch) || isDigit(ch) {
				tokenText = append(tokenText, ch)
				state = Id
			} else {
				l.pos--
				return l.makeToken(Id, string(tokenText))
			}
		case Int2:
			if ch == 't' {
				tokenText = append(tokenText, ch)
				state = Int3
			} else if isLetter(ch) || isDigit(ch) {
				tokenText = append(tokenText, ch)
				state = Id
			} else {
				l.pos--
				return l.makeToken(Id, string(tokenText))
			}
		case Int3:
			if isBlank(ch) {
				return Token{Type: INT_TYPE, Literal: string(tokenText)}
			} else if isLetter(ch) || isDigit(ch) {
				tokenText = append(tokenText, ch)
				state = Id
			} else {
				l.pos--
				return l.makeToken(Id, string(tokenText))
			}
		case IntLiteral:
			if isDigit(ch) {
				tokenText = append(tokenText, ch)
			} else {
				l.pos--
				return Token{Type: INT, Literal: string(tokenText)}
			}
		case GTState:
			if ch == '=' {
				return Token{Type: GE, Literal: ">="}
			} else {
				l.pos--
				return Token{Type: GT, Literal: ">"}
			}
		}
	}
}

func (l *Lexer) makeToken(state DfaState, text string) Token {
	if state == Int3 && text == "int" {
		return Token{Type: INT_TYPE, Literal: text}
	}
	if tok, ok := keywords[text]; ok {
		return Token{Type: tok, Literal: text}
	}
	return Token{Type: IDENT, Literal: text}
}

func main() {
	input := `age >= 45
int age = 40
2+3*5`

	l := New(input)
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
