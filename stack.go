package main

import "errors"

type stack struct {
	elements []rune
}

func newStack() *stack {
	return &stack{make([]rune, 0)}
}

func (s *stack) pop() (rune, error) {
	length := len(s.elements)
	if length == 0 {
		return 0, errors.New(("stack empty"))
	}

	res := s.elements[length-1]
	s.elements = s.elements[:length-1]
	return res, nil
}

func (s *stack) push(v rune) {
	s.elements = append(s.elements, v)
}

func (s *stack) peek() (rune, error) {
	length := len(s.elements)
	if length == 0 {
		return 0, errors.New(("stack empty"))
	}

	return s.elements[length-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.elements) == 0
}
