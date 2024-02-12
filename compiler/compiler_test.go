package compiler

import (
	"testing"
)

func TestCompiler_JitCompile(t *testing.T) {
	runWithCompiler(t, func(t *testing.T, compiler *Compiler) {
		content := []byte("++++++++++++++++++++++++++++++++++++++++++.")
		compiler.JitCompile(content)
	})
}

func runWithCompiler(t *testing.T, test func(t *testing.T, compiler *Compiler)) {
	compiler := NewCompiler(WithMemory(1024))

	test(t, compiler)
}
