package aocutils

import (
	"container/list"
	"errors"
	"fmt"
)

// A stack to store strings only. For a general stack, or other general data structure use other packages like https://github.com/zyedidia/generic
type StringStack struct {
	l list.List
}

// Returns a new initialized string stack. Push and Pop can work right away.
func NewStringStack() *StringStack {
	st := &StringStack{l: *list.New()}
	return st
}

// Returns a new string stack initialized with the values provided. The values are pushed onto the stack from the starting of the values till the end.
func NewPopulatedStringStack(values []string) *StringStack {
	st := NewStringStack()
	for _, v := range values {
		st.Push(v)
	}
	return st
}

// Push a string onto the stack
func (s *StringStack) Push(value string) {
	s.l.PushFront(value)
}

// Remove the top string from the stack and return it. It returns an error if the stack is empty.
func (s *StringStack) Pop() (string, error) {
	if s.Len() == 0 {
		return "", errors.New("Stack underflow error, no element to pop from the stack")
	}

	topElement := s.l.Front()
	topValue := topElement.Value
	s.l.Remove(topElement)

	return fmt.Sprint(topValue), nil
}

// Returns the top string from the stack but does not remove it. It returns an error if the stack is empty.
func (s *StringStack) Peek() (string, error) {
	if s.Len() == 0 {
		return "", errors.New("No elements in stack to peek")
	}

	value := s.l.Front().Value
	return fmt.Sprint(value), nil
}

// Returns the length of the stack
func (s *StringStack) Len() int {
	return s.l.Len()
}

// Returns whether the stack is empty or not
func (s *StringStack) IsEmpty() bool {
	return s.Len() == 0
}

// Implements Stringer
func (s *StringStack) String() string {
	if s.Len() == 0 {
		return "StringStack: empty"
	}
	return fmt.Sprintf("StringStack: %v", s.l)

}
