package shared

import (
	"path/filepath"
)

type Stack[T any] []T

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	var ret T

	if len(*s) == 0 {
		return ret, false
	}

	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return val, true
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func IsBrainfuckFile(name string) bool {
	return filepath.Ext(name) == fileExt
}
