package token

type TokenType string

const (

	// 运算符
	PLUS     = "+" // Plus: +
	MINUS    = "-" // Minus: -
	ASTERISK = "*" // Star: *
	SLASH    = "/" // Slash: /

	GE = ">=" // GE: >=
	GT = ">"  // GT: >
	EQ = "==" // EQ: ==
	LE = "<=" // LE: <=
	LT = "<"  // LT: <

	SEMICOLON = ";" // SEMICOLON: ;
	LPAREN    = "(" // LPAREN: (
	RPAREN    = ")" // RPAREN: )

	ASSIGN = "=" // Assignment: =

	// 标识符和字面量
	IDENT          = "IDENT"          // Identifier: 标识符
	INT_LITERAL    = "INT_LITERAL"    // IntLiteral: 整型字面量
	STRING_LITERAL = "STRING_LITERAL" // StringLiteral: 字符串字面量

	// 数据类型
	INT_TYPE = "INT_TYPE" // Int: int 关键字

	// 关键字
	IF   = "IF"
	ELSE = "ELSE"
)

type Token struct {
	Type    TokenType
	Literal string
}
