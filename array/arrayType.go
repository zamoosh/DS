package array

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Array[T GenericType] struct {
	Items []T
	Count int
}

type AnyArray[T any] struct {
	Items []T
	Count int
}

func NewArrayType[T GenericType](size int) Array[T] {
	return Array[T]{
		Items: make([]T, size),
		Count: 0,
	}
}

func NewArrayTypeAny[T any](size int) AnyArray[T] {
	return AnyArray[T]{
		Items: make([]T, size),
		Count: 0,
	}
}

func (a *Array[T]) Insert(item T) {
	if a.Count == len(a.Items) {
		a.grow()
	}
	a.Items[a.Count] = item
	a.Count++
}

func (a *AnyArray[T]) Insert(item T) {
	if a.Count == len(a.Items) {
		a.grow()
	}
	a.Items[a.Count] = item
	a.Count++
}

func (a *Array[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(a.Items) {
		return errors.New("count is invalid")
	}

	// remove last index
	var NIL T
	if index == a.Count-1 {
		a.Items[index] = NIL
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

func (a *AnyArray[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(a.Items) {
		return errors.New("count is invalid")
	}

	// remove last index
	var NIL T
	if index == a.Count-1 {
		a.Items[index] = NIL
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

func (a *Array[T]) IndexOf(item T) int {
	for i := 0; i < a.Count; i++ {
		if item == a.Items[i] {
			return i
		}
	}
	return -1
}

func (a *AnyArray[T]) IndexOf(item any) int {
	for i := 0; i < a.Count; i++ {
		if reflect.DeepEqual(item, a.Items[i]) {
			return i
		}
	}
	return -1
}

func (a *Array[T]) Max() T {
	if reflect.TypeOf(a.Items[0]).Kind() == reflect.String {
		log.Panicln("can not compare string")
	}

	var item T
	for i := 0; i < a.Count; i++ {
		if a.Items[i] >= item {
			item = a.Items[i]
		}
	}
	return item
}

func (a *Array[T]) grow() {
	if a.Count == len(a.Items) {
		newArr := make([]T, len(a.Items)*2)
		for i := 0; i < a.Count; i++ {
			newArr[i] = a.Items[i]
		}
		a.Items = newArr
	}
}

func (a *AnyArray[T]) grow() {
	if a.Count == len(a.Items) {
		newArr := make([]T, len(a.Items)*2)
		for i := 0; i < a.Count; i++ {
			newArr[i] = a.Items[i]
		}
		a.Items = newArr
	}
}

func (a *Array[T]) Intersect(another *Array[T]) Array[T] {
	commons := NewArrayType[T](3)
	for i := 0; i < a.Count; i++ {
		for j := 0; j < another.Count; j++ {
			if a.Items[i] == another.Items[j] {
				commons.Insert(a.Items[i])
			}
		}
	}

	return commons.removeCommons()
}

func (a *Array[T]) Contains(item T) bool {
	for i := 0; i < a.Count; i++ {
		if a.Items[i] == item {
			return true
		}
	}
	return false
}

func (a *Array[T]) removeCommons() Array[T] {
	unique := NewArrayType[T](3)
	for i := 0; i < a.Count; i++ {
		if !unique.Contains(a.Items[i]) {
			unique.Insert(a.Items[i])
		}
	}
	return unique
}

func (a *Array[T]) Reverse() Array[T] {
	reverse := NewArrayType[T](a.Count)
	for i := a.Count - 1; i >= 0; i-- {
		reverse.Insert(a.Items[i])
	}
	return reverse
}

func (a *AnyArray[T]) Reverse() AnyArray[T] {
	reverse := NewArrayTypeAny[T](a.Count)
	for i := a.Count - 1; i >= 0; i-- {
		reverse.Insert(a.Items[i])
	}
	return reverse
}

func (a *Array[T]) InsertAt(item T, index int) {
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

func (a *AnyArray[T]) InsertAt(item T, index int) {
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

func (a *Array[T]) Print() {
	format := "%v"
	if reflect.TypeOf(a.Items[0]).Kind() == reflect.String {
		format = "\"%v\""
	}

	fmt.Printf("[")
	for i := 0; i < a.Count; i++ {
		fmt.Printf(format, a.Items[i])
		if i != a.Count-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

func (a *AnyArray[T]) Print() {
	format := "%v"
	if reflect.TypeOf(a.Items[0]).Kind() == reflect.String {
		format = "\"%v\""
	}

	fmt.Printf("[")
	for i := 0; i < a.Count; i++ {
		fmt.Printf(format, a.Items[i])
		if i != a.Count-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}
