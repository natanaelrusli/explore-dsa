package main

import (
	"errors"
)

// Stack represents our stack data structure
type Stack struct {
	items []int
}

// NewStack creates and returns a new stack
func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	// Get the last element
	item := s.items[len(s.items)-1]

	// Remove it from the slice
	s.items = s.items[:len(s.items)-1]

	return item, nil
}

// Peek returns the top element without removing it
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}
