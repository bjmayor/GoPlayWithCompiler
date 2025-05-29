package lexer

import (
	"GoPlayWithCompiler/source/simple_calculator/token"
	"unicode"
	"fmt"
)

// DfaState 有限状态机的各种状态
type DfaState int

const (
	Initial DfaState = iota

	If
	Id_if1
	Id_if2
	Else
	Id_else1
	Id_else2
	Id_else3
	Id_else4
	Int
	Id_int1
	Id_int2
	Id_int3
	Id
	GT
	GE

	Assignment

	Plus
	Minus
	Star
	Slash

	SemiColon
	LeftParen
	RightParen

	IntLiteral
)

// SimpleToken 实现 token.Token 接口
type SimpleToken struct {
	Type token.TokenType
	Text string
}

func (t *SimpleToken) GetType() token.TokenType {
	return t.Type
}

func (t *SimpleToken) GetText() string {
	return t.Text
}

// SimpleTokenReader 是一个简单的 Token 流
type SimpleTokenReader struct {
	tokens []*SimpleToken
	pos    int
}

func NewSimpleTokenReader(tokens []*SimpleToken) *SimpleTokenReader {
	return &SimpleTokenReader{
		tokens: tokens,
		pos:    0,
	}
}

func (r *SimpleTokenReader) Read() *SimpleToken {
	if r.pos < len(r.tokens) {
		tok := r.tokens[r.pos]
		r.pos++
		return tok
	}
	return nil
}

func (r *SimpleTokenReader) Peek() *SimpleToken {
	if r.pos < len(r.tokens) {
		return r.tokens[r.pos]
	}
	return nil
}

func (r *SimpleTokenReader) Unread() {
	if r.pos > 0 {
		r.pos--
	}
}

func (r *SimpleTokenReader) Position() int {
	return r.pos
}

func (r *SimpleTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(r.tokens) {
		r.pos = pos
	}
}

// SimpleLexer 是一个简单的手写词法分析器
type SimpleLexer struct {
	tokenText []rune
	tokens    []*SimpleToken
	token     *SimpleToken
}

// NewSimpleLexer 创建一个新的 SimpleLexer
func NewSimpleLexer() *SimpleLexer {
	return &SimpleLexer{}
}

// isAlpha 判断是否为字母
func isAlpha(ch rune) bool {
	return unicode.IsLetter(ch)
}

// isDigit 判断是否为数字
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// isBlank 判断是否为空白字符
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// initToken 有限状态机进入初始状态
func (l *SimpleLexer) initToken(ch rune) DfaState {
	if len(l.tokenText) > 0 {
		l.token.Text = string(l.tokenText)
		l.tokens = append(l.tokens, l.token)
		l.tokenText = []rune{}
		l.token = &SimpleToken{}
	}

	newState := Initial
	if isAlpha(ch) {
		if ch == 'i' {
			newState = Id_int1
		} else {
			newState = Id
		}
		l.token.Type = token.IDENT
		l.tokenText = append(l.tokenText, ch)
	} else if isDigit(ch) {
		newState = IntLiteral
		l.token.Type = token.INT_LITERAL
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '>' {
		newState = GT
		l.token.Type = token.GT
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '+' {
		newState = Plus
		l.token.Type = token.PLUS
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '-' {
		newState = Minus
		l.token.Type = token.MINUS
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '*' {
		newState = Star
		l.token.Type = token.ASTERISK
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '/' {
		newState = Slash
		l.token.Type = token.SLASH
		l.tokenText = append(l.tokenText, ch)
	} else if ch == ';' {
		newState = SemiColon
		l.token.Type = token.SEMICOLON
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '(' {
		newState = LeftParen
		l.token.Type = token.LPAREN
		l.tokenText = append(l.tokenText, ch)
	} else if ch == ')' {
		newState = RightParen
		l.token.Type = token.RPAREN
		l.tokenText = append(l.tokenText, ch)
	} else if ch == '=' {
		newState = Assignment
		l.token.Type = token.ASSIGN
		l.tokenText = append(l.tokenText, ch)
	} else {
		newState = Initial // skip unknown
	}
	return newState
}

// Tokenize 解析字符串，形成 Token 流
func (l *SimpleLexer) Tokenize(code string) *SimpleTokenReader {
	l.tokens = []*SimpleToken{}
	l.tokenText = []rune{}
	l.token = &SimpleToken{}
	state := Initial
	runes := []rune(code)
	length := len(runes)
	var ch rune
	for i := 0; i < length; {
		ch = runes[i]
		fmt.Printf("Debug: i=%d, ch='%c', state=%d\n", i, ch, state)
		switch state {
		case Initial:
			if isBlank(ch) {
				i++ // skip whitespace
			} else {
				state = l.initToken(ch)
				i++
			}
		case Id:
			if isAlpha(ch) || isDigit(ch) {
				l.tokenText = append(l.tokenText, ch)
				i++
			} else {
				state = l.initToken(ch)
			}
		case GT:
			if ch == '=' {
				l.token.Type = token.GE
				state = GE
				l.tokenText = append(l.tokenText, ch)
				i++
			} else {
				state = l.initToken(ch)
			}
		case GE, Assignment, Plus, Minus, Star, Slash, SemiColon, LeftParen, RightParen:
			state = l.initToken(ch)
		case IntLiteral:
			if isDigit(ch) {
				l.tokenText = append(l.tokenText, ch)
				i++
			} else {
				state = l.initToken(ch)
			}
		case Id_int1:
			if ch == 'n' {
				state = Id_int2
				l.tokenText = append(l.tokenText, ch)
				i++
			} else if isAlpha(ch) || isDigit(ch) {
				state = Id
				l.tokenText = append(l.tokenText, ch)
				i++
			} else {
				state = l.initToken(ch)
			}
		case Id_int2:
			if ch == 't' {
				state = Id_int3
				l.tokenText = append(l.tokenText, ch)
				i++
			} else if isAlpha(ch) || isDigit(ch) {
				state = Id
				l.tokenText = append(l.tokenText, ch)
				i++
			} else {
				state = l.initToken(ch)
			}
		case Id_int3:
			if isBlank(ch) {
				l.token.Type = token.INT_TYPE
				state = l.initToken(ch)
				i++
			} else {
				state = Id
				l.tokenText = append(l.tokenText, ch)
				i++
			}
		default:
			i++
		}
	}
	// 处理最后一个 token
	if len(l.tokenText) > 0 {
		l.initToken(0)
	}
	return NewSimpleTokenReader(l.tokens)
}

// Dump 打印所有 Token
func Dump(reader *SimpleTokenReader) {
	println("text\ttype")
	for {
		tok := reader.Read()
		if tok == nil {
			break
		}
		println(tok.Text + "\t\t" + string(tok.Type))
	}
}
