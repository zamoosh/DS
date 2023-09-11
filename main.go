package main

import (
	"ds/array"
)

func main() {
	numbers := array.Array(3)
	numbers.Insert(1)
	numbers.Insert(2)
	numbers.Insert(3)
	// numbers.Insert(4)

	numbers.Print()

	numbers.RemoveAt(2)

	numbers.Print()
}
