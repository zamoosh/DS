package array

import (
	"errors"
	"fmt"
)

type array struct {
	items []interface{}
	count int
}

func Array(size int) array {
	return array{
		items: make([]interface{}, size),
		count: 0,
	}
}

func (a *array) Insert(item int) {
	if a.count == len(a.items) {
		newArr := make([]interface{}, len(a.items)*2)
		for i := 0; i < a.count; i++ {
			newArr[i] = a.items[i]
		}
		a.items = newArr
	}
	a.items[a.count] = item
	a.count++
}

func (a *array) RemoveAt(index int) error {
	if index < 0 || index >= len(a.items) {
		return errors.New("count is invalid")
	}

	// remove last index
	if index == a.count-1 {
		a.items[index] = nil
		a.count--
		return nil
	}

	for i := index; i < a.count; i++ {
		if i+1 == a.count {
			continue
		}
		a.items[i] = a.items[i+1]
	}
	a.count--
	return nil
}

func (a *array) IndexOf(item int) int {
	for i := 0; i < a.count; i++ {
		if item == a.items[i] {
			return i
		}
	}
	return -1
}

func (a *array) Print() {
	fmt.Printf("[")
	for i := 0; i < a.count; i++ {
		fmt.Print(a.items[i])
		if i != a.count-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}
