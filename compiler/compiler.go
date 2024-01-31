package compiler

import "C"
import (
	"bf/shared"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

type Option func(c *Compiler)

func WithMemory(size int) Option {
	return func(c *Compiler) {
		c.memory = size
	}
}

type Compiler struct {
	*shared.Tokenizer
	*shared.Parser

	memory int
}

func NewCompiler(options ...Option) *Compiler {
	c := Compiler{
		Tokenizer: &shared.Tokenizer{},
		Parser:    &shared.Parser{},
		memory:    1024, // default size
	}

	for _, opt := range options {
		opt(&c)
	}

	return &c
}

func (c *Compiler) JitCompile(file string) {
	if filepath.Ext(file) != ".bf" {
		log.Fatal("expecting .bf files only")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	c.Execute(
		c.Compile(
			c.Parse(
				c.Tokenize(content),
			),
		),
	)

}

func (c *Compiler) Compile(ops []shared.Operator) []byte {
	//var code []byte

	for _, op := range ops {
		switch op.Kind {
		case shared.OpPlus:
		case shared.OpMinus:
		case shared.OpLeftArrow:
		case shared.OpRightArrow:
		case shared.OpLeftBracket:
		case shared.OpRightBracket:
		case shared.OpDot:
		case shared.OpComma:
		default:
			log.Printf("unexpected operator encounterd: %c", op.Kind)
			continue
		}
	}

	return nil
}

func (c *Compiler) Execute(instructions []byte) {
	program := mmap(instructions)

	// allocating memory for the compiled program
	memory := (*byte)(unsafe.Pointer(C.malloc(C.size_t(c.memory))))
	// TODO: defer freeing allocated memory

	println(*memory)
	// executing in-memory program
	program(memory)

	println(*memory)
}
