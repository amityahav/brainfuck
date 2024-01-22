package main

import (
	"fmt"
	"log"
)

type Interpreter struct {
	memory []byte
	head   int

	pc int

	tokenize func([]byte) []Token
	parse    func([]Token) []Operator
}

type InterpreterOption func(i *Interpreter)

func WithMemory(size int) InterpreterOption {
	return func(i *Interpreter) {
		i.memory = make([]byte, size)
	}
}

func NewInterpreter(tokenize func([]byte) []Token,
	parse func([]Token) []Operator,
	options ...InterpreterOption) *Interpreter {
	i := Interpreter{
		memory:   make([]byte, 1024), // default size
		tokenize: tokenize,
		parse:    parse,
	}

	for _, opt := range options {
		opt(&i)
	}

	return &i
}

func (i *Interpreter) Interpret(content []byte) {
	// tokenize the source code
	tokens := i.tokenize(content)

	// parse tokens to intermediate representations
	ops := i.parse(tokens)

	// execute operations
	for i.pc < len(ops) {
		op := ops[i.pc]

		switch op.kind {
		case OpPlus:
			// handle overflows
			i.memory[i.head] += byte(op.Operand)
			i.pc++
		case OpMinus:
			// handle underflow
			i.memory[i.head] -= byte(op.Operand)
			i.pc++
		case OpRightArrow:
			// handle out of bounds
			i.head += op.Operand
			i.pc++
		case OpLeftArrow:
			// handle out of bounds
			i.head -= op.Operand
			i.pc++
		case OpRightBracket:
			if i.memory[i.head] != 0 {
				i.pc = op.Operand
				continue
			}

			i.pc++
		case OpLeftBracket:
			if i.memory[i.head] == 0 {
				i.pc = op.Operand
				continue
			}

			i.pc++
		case OpDot:
			fmt.Print(i.memory[i.head])
			i.pc++
		case OpComma:
			i.pc++
		default:
			log.Printf("unexpected operator encounterd: %s", string(op.kind))
			continue
		}
	}
}
