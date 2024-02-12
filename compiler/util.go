package compiler

import (
	"golang.org/x/sys/unix"
	"os"
	"unsafe"
)

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

	code, err := unix.Mmap(int(f.Fd()), 0, 1024, unix.PROT_EXEC, unix.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	codePtr := &code

	return *(*func(pointer *byte))(unsafe.Pointer(&codePtr))
}
