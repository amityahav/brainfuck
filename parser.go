package main

import (
	"log"
	"math"
)

type Operator struct {
	kind    OpKind
	Operand int
}

const placeHolder int = math.MaxInt

type Parser struct {
	// stack is used to validate brackets
	stack Stack
}

func (p *Parser) Parse(tokens []Token) []Operator {
	operators := make([]Operator, len(tokens))

	for _, token := range tokens {
		var op Operator

		switch token.Kind() {
		case OpPlus, OpMinus, OpLeftArrow, OpRightArrow, OpDot, OpComma:
			op = Operator{
				kind:    token.Kind(),
				Operand: len(token),
			}
		case OpLeftBracket:
			op = Operator{
				kind:    token.Kind(),
				Operand: placeHolder,
			}

			p.stack.Push(len(operators))
		case OpRightBracket:
			op = Operator{
				kind:    token.Kind(),
				Operand: placeHolder,
			}

			pos, ok := p.stack.Pop()
			if !ok || operators[pos].kind != OpLeftBracket {
				log.Fatal("parser: unbalanced brackets")
			}

			operators[pos].Operand = len(operators)
			op.Operand = pos
		default:
			log.Printf("unexpected operator encounterd: %s", string(token.Kind()))
			continue
		}

		operators = append(operators, op)
	}

	if p.stack.Size() > 0 {
		log.Fatal("parser: unbalanced brackets")
	}

	return operators
}
