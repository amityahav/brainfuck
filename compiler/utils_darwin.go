package compiler

import (
	"os"
	"syscall"
	"unsafe"
)

var writeSyscallOpcode = []byte{0x04, 0x00, 0x00, 0x02}

func mmap(instructions []byte) func(pointer *byte) {
	// Apple disallows mmaping with WRITE and EXEC together for security reasons so as a workaround
	// im writing the code to a file and then mmap it as executable
	f, err := os.OpenFile("tmp", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	defer os.Remove("tmp")

	_, err = f.Write(instructions)
	if err != nil {
		panic(err)
	}

	code, err := syscall.Mmap(int(f.Fd()), 0, len(instructions), syscall.PROT_EXEC, syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	codePtr := &code

	return *(*func(pointer *byte))(unsafe.Pointer(&codePtr))
}
