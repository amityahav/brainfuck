package interpreter

import (
	"bf/shared"
	"fmt"
	"log"
)

type Option func(i *Interpreter)

func WithMemory(size int) Option {
	return func(i *Interpreter) {
		i.memory = size
	}
}

type Interpreter struct {
	*shared.Tokenizer
	*shared.Parser

	memory int
}

func NewInterpreter(options ...Option) *Interpreter {
	i := Interpreter{
		Tokenizer: &shared.Tokenizer{},
		Parser:    &shared.Parser{},
		memory:    1024, // default size
	}

	for _, opt := range options {
		opt(&i)
	}

	return &i
}

func (i *Interpreter) Interpret(content []byte) {
	i.Execute(
		i.Parse(
			i.Tokenize(content),
		),
	)
}

func (i *Interpreter) Execute(ops []shared.Operator) {
	var (
		head int
		pc   int
	)

	memory := make([]byte, i.memory)

	for pc < len(ops) {
		if head < 0 || head >= i.memory {
			log.Fatal("head points to invalid memory address")
		}

		op := ops[pc]

		switch op.Kind {
		case shared.OpPlus:
			memory[head] += byte(op.Operand)
			pc++
		case shared.OpMinus:
			memory[head] -= byte(op.Operand)
			pc++
		case shared.OpRightArrow:
			head += op.Operand
			pc++
		case shared.OpLeftArrow:
			head -= op.Operand
			pc++
		case shared.OpRightBracket:
			if memory[head] != 0 {
				pc = op.Operand
				continue
			}

			pc++
		case shared.OpLeftBracket:
			if memory[head] == 0 {
				pc = op.Operand
				continue
			}

			pc++
		case shared.OpDot:
			fmt.Printf("%c", memory[head])
			pc++
		default:
			log.Printf("unexpected operator encounterd: %c", op.Kind)
			continue
		}
	}
}
