package main

import (
	"ds/linkList"
)

func main() {
	numbers := linkList.LinkList()
	numbers.AddFirst(30)
	numbers.AddLast(10)
	numbers.AddLast(20)
	numbers.AddFirst(40)
	numbers.AddFirst(50)

	numbers.Print()

	// numbers.DeleteLast()
	// numbers.DeleteFirst()
	//
	// numbers.Print()

	println(numbers.IndexOf(10))
}
