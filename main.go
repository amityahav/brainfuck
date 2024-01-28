package main

import (
	"bf/compiler"
	"bf/interpreter"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough arguments")
	}

	mode := os.Args[1]
	bfile := os.Args[2]

	switch mode {
	case "interpret":
		i := interpreter.NewInterpreter(
			interpreter.WithMemory(8192),
		)

		i.Interpret(bfile)
	case "compile":
		c := compiler.NewCompiler(
			compiler.WithMemory(8192),
		)

		c.JitCompile(bfile)
	}

}
