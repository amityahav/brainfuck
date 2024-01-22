package main

import (
	"log"
	"os"
)

func main() {
	bfile := os.Args[1]

	content, err := os.ReadFile(bfile)
	if err != nil {
		log.Fatal(err)
	}

	interpreter := NewInterpreter(
		Tokenzier{}.Tokenize,
		Parser{}.Parse,
		WithMemory(8192),
	)

	interpreter.Interpret(content)
}
