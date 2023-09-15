package main

import (
	"ds/stack"
)

func main() {
	numbers := stack.NewStack()
	numbers.Push(10)
	numbers.Push(20)
	numbers.Push(30)
	numbers.Push(40)
	numbers.Push(50)

	numbers.Print()
	
	println(numbers.Pop())
	
	numbers.Print()
}
