package stack

import (
	"ds/array"
	"log"
	"reflect"
)

type stack[T array.GenericType] struct {
	content array.Array[T]
}

func NewStack[T array.GenericType]() stack[T] {
	return stack[T]{
		content: array.NewArrayType[T](5),
	}
}

func (s *stack[T]) Push(item T) {
	s.content.Insert(item)
}

func (s *stack[T]) Pop() T {
	item := s.Peek()

	err := s.content.RemoveAt(s.content.Count - 1)
	if err != nil {
		println("bad index")
	}
	return item
}

func (s *stack[T]) Peek() T {
	if s.content.Count == 0 {
		log.Panic("stack is empty")
	}
	item := s.content.Items[s.content.Count-1]
	return item
}

func (s *stack[T]) IsEmpty() bool {
	if s.content.Count != 0 {
		return false
	}
	return true
}

func (s *stack[T]) IsStable(exp interface{}) bool {
	if reflect.TypeOf(exp).Kind() != reflect.String {
		log.Panicln("exp must be string")
	}

	switch ch := exp.(type) {
	case string:
		for i := range ch {
			switch string(ch[i]) {
			case "(", "[", "<":
				s.Push(T(ch[i]))
			case ")", "]", ">":
				if s.IsEmpty() {
					return false
				}
				s.Pop()
			}
		}
	default:
		log.Panicln("Unsupported type")
	}
	if s.IsEmpty() {
		return true
	}
	return false
}

func (s *stack[T]) Print() {
	s.content.Print()
}
