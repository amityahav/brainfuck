package compiler

import (
	"bf/shared"
	"log"
)

// back patching
type bp struct {
	// current instruction position
	cp int32
	// next instruction position
	np int32
	// placeholder position
	pp int32
}

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
	stack  shared.Stack[bp]
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

func (c *Compiler) JitCompile(content []byte) {
	c.Execute(
		c.CompileX86(
			c.Parse(
				c.Tokenize(content),
			),
		),
	)

}

// CompileX86 compiles to x86-64 machine code
func (c *Compiler) CompileX86(ops []shared.Operator) []byte {
	var (
		code []byte
	)

	for _, op := range ops {
		switch op.Kind {
		case shared.OpPlus:
			// ADD BYTE[rax], operand
			code = append(code, 0x80, 0x00, byte(op.Operand))
		case shared.OpMinus:
			// SUB BYTE[rax], operand
			code = append(code, 0x80, 0x28, byte(op.Operand))
		case shared.OpLeftArrow:
			// SUB rax, operand
			code = append(code, 0x48, 0x83, 0xE8, byte(op.Operand))
		case shared.OpRightArrow:
			// ADD rax, operand
			code = append(code, 0x48, 0x83, 0xC0, byte(op.Operand))
		case shared.OpLeftBracket:
			backPatch := bp{
				cp: int32(len(code)),
			}

			// MOV bl, BYTE PTR[rax]
			// cmp bl, 0
			// je <matching right bracket addr>
			code = append(code,
				0x8A, 0x18,
				0x80, 0xFB, 00,
				0x0F, 0x84)

			backPatch.pp = int32(len(code))
			// last 4 bytes are placeholders for the addr that will be back patched once we encounter
			// the matching right bracket
			code = append(code, 0x00, 0x00, 0x00, 0x00)

			backPatch.np = int32(len(code))

			c.stack.Push(backPatch)
		case shared.OpRightBracket:
			backPatch, ok := c.stack.Pop()
			if !ok {
				log.Fatal("compiler: unbalanced brackets")
			}

			// back patch
			cp := int32(len(code))
			relative := int32toLittleEndian(cp - backPatch.np)
			for i := 0; i < 4; i++ {
				code[int(backPatch.pp)+i] = relative[i]
			}

			// MOV bl, BYTE PTR[rax]
			// cmp bl, 0
			// jne <matching left bracket addr>
			code = append(code,
				0x8A, 0x18,
				0x80, 0xFB, 00,
				0x0F, 0x85)

			np := int32(len(code) + 4)
			relative = int32toLittleEndian(backPatch.cp - np)
			code = append(code, relative...)
		case shared.OpDot:
			// MOV r9, rax 			 ; saving rax (memory head) because it is needed for system call
			// MOV rax, <write_syscall_opcode>
			// MOV rdi,0x1			 ; stdout
			// MOV rsi,r9         	 ; r9's value is the memory head pointer
			// MOV rdx,0x1	     	 ; count
			// syscall
			// MOV rax, r9           ; restoring rax
			code = append(code,
				0x49, 0x89, 0xc1,
				0x48, 0xC7, 0xC0)

			code = append(code, writeSyscallOpcode...)

			code = append(code,
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

	//for _, b := range code {
	//	fmt.Printf("%02x ", b)
	//}

	return code
}

func (c *Compiler) Execute(instructions []byte) {
	program := mmap(instructions)

	memory := make([]byte, c.memory)
	// executing in-memory program
	program(&memory[0])
}
