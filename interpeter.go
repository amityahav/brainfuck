package main

type Interpreter struct {
	memory []byte
	ip     int

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

	// execute operations one-by-one
	for _, op := range ops {
		// TODO
	}
}
