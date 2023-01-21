package main

import (
	"fmt"
	"os"
)

func main() {
	src, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens := lex(string(src))
	for _, token := range tokens {
		fmt.Println(token.value)
	}

	ast, _ := parse(tokens, 0)
	fmt.Println(ast.pretty())
}
