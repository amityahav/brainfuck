package compiler

import "golang.org/x/sys/unix"

var writeSyscallOpcode = []byte{0x01, 0x00, 0x00, 0x00}

func mmap(instructions []byte) func(pointer *byte) {
	code, err := unix.Mmap(-1, 0, 1024, unix.PROT_EXEC|unix.PROT_WRITE|unix.PROT_READ, unix.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	copy(code)

	codePtr := &code

	return *(*func(pointer *byte))(unsafe.Pointer(&codePtr))
}
