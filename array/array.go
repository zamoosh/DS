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

func (a *array) Max() int {
	item := 0
	for i := 0; i < a.count; i++ {
		if a.items[i].(int) >= item {
			item = a.items[i].(int)
		}
	}
	return item
}

func (a *array) Intersect(another *array) array {
	commons := Array(3)
	for i := 0; i < a.count; i++ {
		for j := 0; j < another.count; j++ {
			if a.items[i] == another.items[j] {
				commons.Insert(a.items[i].(int))
			}
		}
	}
	unique := Array(3)
	for i := 0; i < commons.count; i++ {
		if !unique.Contains(commons.items[i].(int)) {
			unique.Insert(commons.items[i].(int))
		}
	}
	return unique
}

func (a *array) Contains(item int) bool {
	for i := 0; i < a.count; i++ {
		if a.items[i] == item {
			return true
		}
	}
	return false
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

// [1, 2, 2, 5]
// [1, 2, 5]
