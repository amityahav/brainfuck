package compiler

import (
	"bf/shared"
	"log"
	"os"
	"path/filepath"
)

type Option func(c *Compiler)

func WithMemory(size int) Option {
	return func(c *Compiler) {
		c.memory = make([]byte, size)
	}
}

type Compiler struct {
	*shared.Tokenizer
	*shared.Parser

	memory []byte
}

func NewCompiler(options ...Option) *Compiler {
	c := Compiler{
		Tokenizer: &shared.Tokenizer{},
		Parser:    &shared.Parser{},
		memory:    make([]byte, 1024), // default size
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

func (c *Compiler) Compile(ops []shared.Operator) func(memory []byte) {
	return nil
}

func (c *Compiler) Execute(code func(memory []byte)) {
	code(c.memory)
}
