package array

import (
	"errors"
	"fmt"
	"log"
)

type Array struct {
	Items []interface{}
	Count int
}

func NewArrayType(size int) Array {
	return Array{
		Items: make([]interface{}, size),
		Count: 0,
	}
}

func (a *Array) Insert(item int) {
	if a.Count == len(a.Items) {
		a.grow()
	}
	a.Items[a.Count] = item
	a.Count++
}

func (a *Array) RemoveAt(index int) error {
	if index < 0 || index >= len(a.Items) {
		return errors.New("Count is invalid")
	}

	// remove last index
	if index == a.Count-1 {
		a.Items[index] = nil
		a.Count--
		return nil
	}

	for i := index; i < a.Count; i++ {
		if i+1 == a.Count {
			continue
		}
		a.Items[i] = a.Items[i+1]
	}
	a.Count--
	return nil
}

func (a *Array) IndexOf(item int) int {
	for i := 0; i < a.Count; i++ {
		if item == a.Items[i] {
			return i
		}
	}
	return -1
}

func (a *Array) Max() int {
	item := 0
	for i := 0; i < a.Count; i++ {
		if a.Items[i].(int) >= item {
			item = a.Items[i].(int)
		}
	}
	return item
}

func (a *Array) grow() {
	if a.Count == len(a.Items) {
		newArr := make([]interface{}, len(a.Items)*2)
		for i := 0; i < a.Count; i++ {
			newArr[i] = a.Items[i]
		}
		a.Items = newArr
	}
}

func (a *Array) Intersect(another *Array) array {
	commons := NewArray(3)
	for i := 0; i < a.Count; i++ {
		for j := 0; j < another.Count; j++ {
			if a.Items[i] == another.Items[j] {
				commons.Insert(a.Items[i].(int))
			}
		}
	}

	return commons.removeCommons()
}

func (a *Array) Contains(item int) bool {
	for i := 0; i < a.Count; i++ {
		if a.Items[i] == item {
			return true
		}
	}
	return false
}

func (a *Array) removeCommons() array {
	unique := NewArray(3)
	for i := 0; i < a.Count; i++ {
		if !unique.Contains(a.Items[i].(int)) {
			unique.Insert(a.Items[i].(int))
		}
	}
	return unique
}

func (a *Array) Reverse() array {
	reverse := NewArray(a.Count)
	for i := a.Count - 1; i >= 0; i-- {
		reverse.Insert(a.Items[i].(int))
	}
	return reverse
}

func (a *Array) InsertAt(item int, index int) {
	if index == a.Count {
		a.Insert(item)
	} else if index == 0 {
		if a.Count == len(a.Items) {
			a.grow()
		}
		for i := a.Count - 1; i >= 0; i-- {
			a.Items[i+1] = a.Items[i]
		}
		a.Items[index] = item
		a.Count++
	} else if index > 0 && index < a.Count {
		if a.Count == len(a.Items) {
			a.grow()
		}
		for i := a.Count - 1; i >= index; i-- {
			a.Items[i+1] = a.Items[i]
		}
		a.Items[index] = item
		a.Count++
	} else {
		log.Fatalln("invalid index")
	}
}

func (a *Array) Print() {
	fmt.Printf("[")
	for i := 0; i < a.Count; i++ {
		fmt.Print(a.Items[i])
		if i != a.Count-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

// [1, 2, 2, 5] -> InsertAt(0, 2): index=2
// [1, 2, 0, 2, 5]
