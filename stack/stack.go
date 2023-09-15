package stack

import (
	"ds/array"
	"log"
)

type stack struct {
	content array.Array
}

func NewStack() stack {
	return stack{
		content: array.NewArrayType(5),
	}
}

func (s *stack) Push(item int) {
	s.content.Insert(item)
}

func (s *stack) Pop() int {
	item := s.Peek()

	err := s.content.RemoveAt(s.content.Count - 1)
	if err != nil {
		println("bad index")
	}
	return item
}

func (s *stack) Peek() int {
	if s.content.Count == 0 {
		log.Panic("stack is empty")
	}
	item := s.content.Items[s.content.Count-1].(int)
	return item
}

func (s *stack) IsEmpty() bool {
	if s.content.Count != 0 {
		return false
	}
	return true
}

func (s *stack) Print() {
	s.content.Print()
}

// [1, 2, 3, 4, 5]
