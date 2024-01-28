package interpreter

import (
	"bf/shared"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Option func(i *Interpreter)

func WithMemory(size int) Option {
	return func(i *Interpreter) {
		i.memory = make([]byte, size)
	}
}

type Interpreter struct {
	*shared.Tokenizer
	*shared.Parser

	memory []byte
	head   int

	pc int
}

func NewInterpreter(options ...Option) *Interpreter {
	i := Interpreter{
		Tokenizer: &shared.Tokenizer{},
		Parser:    &shared.Parser{},
		memory:    make([]byte, 1024), // default size
	}

	for _, opt := range options {
		opt(&i)
	}

	return &i
}

func (i *Interpreter) Interpret(file string) {
	if filepath.Ext(file) != ".bf" {
		log.Fatal("expecting .bf files only")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	i.Execute(
		i.Parse(
			i.Tokenize(content),
		),
	)
}

func (i *Interpreter) Execute(ops []shared.Operator) {
	for i.pc < len(ops) {
		if i.head < 0 || i.head >= len(i.memory) {
			log.Fatal("head points to invalid memory address")
		}

		op := ops[i.pc]

		switch op.Kind {
		case shared.OpPlus:
			i.memory[i.head] += byte(op.Operand)
			i.pc++
		case shared.OpMinus:
			i.memory[i.head] -= byte(op.Operand)
			i.pc++
		case shared.OpRightArrow:
			i.head += op.Operand
			i.pc++
		case shared.OpLeftArrow:
			i.head -= op.Operand
			i.pc++
		case shared.OpRightBracket:
			if i.memory[i.head] != 0 {
				i.pc = op.Operand
				continue
			}

			i.pc++
		case shared.OpLeftBracket:
			if i.memory[i.head] == 0 {
				i.pc = op.Operand
				continue
			}

			i.pc++
		case shared.OpDot:
			fmt.Printf("%c", i.memory[i.head])
			i.pc++
		case shared.OpComma: // TODO
			i.pc++
		default:
			log.Printf("unexpected operator encounterd: %s", string(op.Kind))
			continue
		}
	}
}
