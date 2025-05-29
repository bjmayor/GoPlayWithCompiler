// https://zupzup.org/go-ast-traversal/
package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello"
	world := "World"
	words := []string{hello, world}
	SayHello(words)
}

// SayHello says Hello
func SayHello(words []string) {
	fmt.Println(joinStrings(words))
}

// joinStrings joins strings
func joinStrings(words []string) string {
	return strings.Join(words, ", ")
}
