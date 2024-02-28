package compiler

func int32toLittleEndian(x int32) []byte {
	var b [4]byte

	b[0] = byte(x)
	x = x >> 8
	b[1] = byte(x)
	x = x >> 8
	b[2] = byte(x)
	x = x >> 8
	b[3] = byte(x)

	return b[:]
}
