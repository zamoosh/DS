package array

import (
	"errors"
	"fmt"
	"log"
)

type array struct {
	items []interface{}
	count int
}

func NewArray(size int) array {
	return array{
		items: make([]interface{}, size),
		count: 0,
	}
}

func (a *array) Insert(item int) {
	if a.count == len(a.items) {
		a.grow()
	}
	a.items[a.count] = item
	a.count++
}

func (a *array) RemoveAt(index int) error {
	if index < 0 || index >= len(a.items) {
		return errors.New("Count is invalid")
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

func (a *array) grow() {
	if a.count == len(a.items) {
		newArr := make([]interface{}, len(a.items)*2)
		for i := 0; i < a.count; i++ {
			newArr[i] = a.items[i]
		}
		a.items = newArr
	}
}

func (a *array) Intersect(another *array) array {
	commons := NewArray(3)
	for i := 0; i < a.count; i++ {
		for j := 0; j < another.count; j++ {
			if a.items[i] == another.items[j] {
				commons.Insert(a.items[i].(int))
			}
		}
	}

	return commons.removeCommons()
}

func (a *array) Contains(item int) bool {
	for i := 0; i < a.count; i++ {
		if a.items[i] == item {
			return true
		}
	}
	return false
}

func (a *array) removeCommons() array {
	unique := NewArray(3)
	for i := 0; i < a.count; i++ {
		if !unique.Contains(a.items[i].(int)) {
			unique.Insert(a.items[i].(int))
		}
	}
	return unique
}

func (a *array) Reverse() array {
	reverse := NewArray(a.count)
	for i := a.count - 1; i >= 0; i-- {
		reverse.Insert(a.items[i].(int))
	}
	return reverse
}

func (a *array) InsertAt(item int, index int) {
	if index == a.count {
		a.Insert(item)
	} else if index == 0 {
		if a.count == len(a.items) {
			a.grow()
		}
		for i := a.count - 1; i >= 0; i-- {
			a.items[i+1] = a.items[i]
		}
		a.items[index] = item
		a.count++
	} else if index > 0 && index < a.count {
		if a.count == len(a.items) {
			a.grow()
		}
		for i := a.count - 1; i >= index; i-- {
			a.items[i+1] = a.items[i]
		}
		a.items[index] = item
		a.count++
	} else {
		log.Fatalln("invalid index")
	}
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

// [1, 2, 2, 5] -> InsertAt(0, 2): index=2
// [1, 2, 0, 2, 5]
