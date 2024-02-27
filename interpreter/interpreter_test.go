package interpreter

import "testing"

func TestInterpreter_Interpret(t *testing.T) {
	runWithInterpreter(t, func(t *testing.T, interpreter *Interpreter) {
		content := []byte("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
		interpreter.Interpret(content)
	})
}

func runWithInterpreter(t *testing.T, test func(t *testing.T, interpreter *Interpreter)) {
	interpreter := NewInterpreter(WithMemory(1024))

	test(t, interpreter)
}
