package shared

import (
	"path/filepath"
)

type Stack []int

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return val, true
}

func (s *Stack) Size() int {
	return len(*s)
}

func IsBrainfuckFile(name string) bool {
	return filepath.Ext(name) == fileExt
}
