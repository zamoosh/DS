package array

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
	"reflect"
)

type GenericType interface {
	constraints.Ordered
}

type array[T GenericType] struct {
	items []T
	count int
}

func NewArray[T GenericType](size int) array[T] {
	return array[T]{
		items: make([]T, size),
		count: 0,
	}
}

func (a *array[T]) Insert(item T) {
	if a.count == len(a.items) {
		a.grow()
	}
	a.items[a.count] = item
	a.count++
}

func (a *array[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(a.items) {
		return errors.New("Count is invalid")
	}

	// remove last index
	var NIL T
	if index == a.count-1 {
		a.items[index] = NIL
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

func (a *array[T]) IndexOf(item T) int {
	for i := 0; i < a.count; i++ {
		if item == a.items[i] {
			return i
		}
	}
	return -1
}

func (a *array[T]) Max() T {
	if reflect.TypeOf(a.items[0]).String() == "string" {
		log.Panicln("can not compare string")
	}

	var item T
	for i := 0; i < a.count; i++ {
		if a.items[i] >= item {
			item = a.items[i]
		}
	}
	return item
}

func (a *array[T]) grow() {
	if a.count == len(a.items) {
		newArr := make([]T, len(a.items)*2)
		for i := 0; i < a.count; i++ {
			newArr[i] = a.items[i]
		}
		a.items = newArr
	}
}

func (a *array[T]) Intersect(another *array[T]) array[T] {
	commons := NewArray[T](3)
	for i := 0; i < a.count; i++ {
		for j := 0; j < another.count; j++ {
			if a.items[i] == another.items[j] {
				commons.Insert(a.items[i])
			}
		}
	}

	return commons.removeCommons()
}

func (a *array[T]) Contains(item T) bool {
	for i := 0; i < a.count; i++ {
		if a.items[i] == item {
			return true
		}
	}
	return false
}

func (a *array[T]) removeCommons() array[T] {
	unique := NewArray[T](3)
	for i := 0; i < a.count; i++ {
		if !unique.Contains(a.items[i]) {
			unique.Insert(a.items[i])
		}
	}
	return unique
}

func (a *array[T]) Reverse() array[T] {
	reverse := NewArray[T](a.count)
	for i := a.count - 1; i >= 0; i-- {
		reverse.Insert(a.items[i])
	}
	return reverse
}

func (a *array[T]) InsertAt(item T, index int) {
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

func (a *array[T]) Print() {
	format := "%v"
	if reflect.TypeOf(a.items[0]).String() == "string" {
		format = "\"%v\""
	}

	fmt.Printf("[")
	for i := 0; i < a.count; i++ {
		fmt.Printf(format, a.items[i])
		if i != a.count-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}
