package compiler

import (
	"syscall"
	"unsafe"
)

var writeSyscallOpcode = []byte{0x01, 0x00, 0x00, 0x00}

func mmap(instructions []byte) func(pointer *byte) {
	code, err := syscall.Mmap(-1, 0, len(instructions), syscall.PROT_EXEC|syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_PRIVATE|syscall.MAP_ANON)
	if err != nil {
		panic(err)
	}

	copy(code, instructions)

	codePtr := &code

	return *(*func(pointer *byte))(unsafe.Pointer(&codePtr))
}
