package main

import (
	"bf/compiler"
	"bf/interpreter"
	"bf/shared"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough arguments")
	}

	mode := os.Args[1]
	bfile := os.Args[2]

	if !shared.IsBrainfuckFile(bfile) {
		log.Fatal("expecting .bf files only")
	}

	content, err := os.ReadFile(bfile)
	if err != nil {
		log.Fatal(err)
	}

	switch mode {
	case "interpret":
		i := interpreter.NewInterpreter(
			interpreter.WithMemory(8192),
		)

		i.Interpret(content)
	case "compile":
		c := compiler.NewCompiler(
			compiler.WithMemory(8192),
		)

		c.JitCompile(content)
	}

}
