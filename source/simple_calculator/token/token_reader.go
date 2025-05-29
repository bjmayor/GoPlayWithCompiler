package token

// TokenReader 表示一个 Token 流，由 Lexer 生成，Parser 可以从中获取 Token。
type TokenReader interface {
	// Read 返回下一个 Token，并从流中取出。如果流已空，返回 nil。
	Read() *Token

	// Peek 返回下一个 Token，但不取出。如果流已空，返回 nil。
	Peek() *Token

	// Unread 回退一步，恢复上一个 Token。
	Unread()

	// Position 返回当前读取位置。
	Position() int

	// SetPosition 设置当前读取位置。
	SetPosition(pos int)
}
