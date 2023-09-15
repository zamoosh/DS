package main

import (
	"ds/linkList"
)

func main() {
	numbers := linkList.LinkList()
	numbers.AddLast(10)
	numbers.AddLast(20)
	numbers.AddLast(30)
	numbers.AddLast(40)
	numbers.AddLast(50)

	numbers.Print()
	println(numbers.KthNodeFromEnd(1))
}
