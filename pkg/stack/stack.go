package stack

import (
	"errors"
)

var (
	ErrStackIsEmpty = errors.New("stack is empty")
)

// A simple stack of any type
type Stack[T any] struct {
	items []T
}

// Pushes a new item on top of the stack
func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

// Removes the item on top of the stack
func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

// Returns the item on top of the stack
func (s *Stack[T]) Top() (*T, error) {
	if s.IsEmpty() {
		return nil, ErrStackIsEmpty
	}

	return &s.items[len(s.items)-1], nil
}

// True if there are no stacked items
func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}
