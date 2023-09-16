package stack

import (
	"ds/array"
	"log"
	"reflect"
)

type stack struct {
	content array.AnyArray[interface{}]
}

func NewStack() stack {
	return stack{
		content: array.NewArrayTypeAny[interface{}](5),
	}
}

func (s *stack) Push(item interface{}) {
	s.content.Insert(item)
}

func (s *stack) Pop() interface{} {
	item := s.Peek()

	err := s.content.RemoveAt(s.content.Count - 1)
	if err != nil {
		println("bad index")
	}
	return item
}

func (s *stack) Peek() interface{} {
	if s.content.Count == 0 {
		log.Panic("stack is empty")
	}
	item := s.content.Items[s.content.Count-1]
	return item
}

func (s *stack) IsEmpty() bool {
	if s.content.Count != 0 {
		return false
	}
	return true
}

func (s *stack) IsStable(exp any) bool {
	if reflect.TypeOf(exp).Kind() != reflect.String {
		log.Panicln("exp must be string")
	}

	switch conExp := exp.(type) {
	case string:
		for i := range conExp {
			char := conExp[i]
			switch string(char) {
			case "(", "[", "<":
				s.Push(char)
			case ")", "]", ">":
				if s.IsEmpty() {
					return false
				}
				popped := s.Pop().(byte)
				switch string(char) {
				case ")":
					if popped != '(' {
						return false
					}
				case "]":
					if popped != '[' {
						return false
					}
				case ">":
					if popped != '<' {
						return false
					}
				}
			}
		}
	default:
		log.Panicln("Unsupported type")
	}
	return s.IsEmpty()
}

func (s *stack) Print() {
	s.content.Print()
}
