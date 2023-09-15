package main

import (
	"ds/linkList"
	"fmt"
)

func main() {
	numbers := linkList.LinkList()
	numbers.AddLast(30)
	// numbers.AddLast(10)
	// numbers.AddLast(20)
	// numbers.AddFirst(40)
	// numbers.AddFirst(50)

	numbers.Print()
	println(numbers.Size())

	// numbers.DeleteLast()
	// numbers.DeleteFirst()
	//
	// numbers.Print()

	numbers.Print()
	println(numbers.Size())

	println("____________________")
	arr := numbers.ToArray()
	fmt.Println(arr)
	fmt.Println(arr[0])

}
