package compiler

import "C"
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

func (c *Compiler) Compile(ops []shared.Operator) []byte {
	var code []byte

	for _, op := range ops {
		switch op.Kind {
		case shared.OpPlus:
			// NOP
			// ADD BYTE[rax], operand
			code = append(code, 0x90, 0x80, 0x00, byte(op.Operand))
		case shared.OpMinus:
			// NOP
			// SUB BYTE[rax], operand
			code = append(code, 0x90, 0x80, 0x28, byte(op.Operand))
		case shared.OpLeftArrow:
			// SUB rax, operand
			code = append(code, 0x48, 0x83, 0xE8, byte(op.Operand))
		case shared.OpRightArrow:
			// ADD rax, operand
			code = append(code, 0x48, 0x83, 0xC0, byte(op.Operand))
		case shared.OpLeftBracket:
		case shared.OpRightBracket:
		case shared.OpDot:
			// MOV r9, rax 			 ; saving rax because it is needed for system call
			// MOV rax,0x2000004	 ; number of write syscall in mac
			// MOV rdi,0x1			 ; stdout
			// MOV rsi,r9         	 ; r9's value is the memory head pointer
			// MOV rdx,0x1	     	 ; count
			// syscall
			// MOV rax, r9           ; restoring rax
			code = append(code,
				0x49, 0x89, 0xc1,
				0x48, 0xC7, 0xC0, 0x04, 0x00, 0x00, 0x02,
				0x48, 0xC7, 0xC7, 0x01, 0x00, 0x00, 0x00,
				0x4C, 0x89, 0xCE,
				0x48, 0xc7, 0xc2, 0x01, 0x00, 0x00, 0x00,
				0x0f, 0x05,
				0x4C, 0x89, 0xC8)
		default:
			log.Printf("unexpected operator encounterd: %c", op.Kind)
			continue
		}
	}

	code = append(code, 0xC3) // RET

	return nil
}

func (c *Compiler) Execute(instructions []byte) {
	program := mmap(instructions)

	// executing in-memory program
	program(&c.memory[0])
}
