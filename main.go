package main

import (
	"ds/array"
)

func main() {
	numbers := array.Array(3)
	numbers.Insert(1)
	numbers.Insert(2)
	numbers.Insert(3)
	numbers.Insert(4)
	numbers.Insert(5)
	numbers.Insert(6)
	// numbers.Insert(5)
	// numbers.Insert(5)

	numbers.Print()
	// another := array.Array(3)
	// another.Insert(5)
	// another.Insert(4)
	//
	// c := numbers.Intersect(&another)
	// c.Print()
	//
	// reverse := numbers.Reverse()
	// reverse.Print()
	numbers.InsertAt(0, 4)
	numbers.Print()
}
