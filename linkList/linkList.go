package linkList

import (
	"fmt"
	"log"
)

type node struct {
	next  *node
	value int
}

type linkList struct {
	first *node
	last  *node
	size  int
}

func LinkList() linkList {
	l := linkList{first: nil, last: nil}
	return l
}

func (l *linkList) isEmpty() bool {
	if l.first != nil {
		return false
	}
	return true
}

func (l *linkList) AddFirst(item int) {
	n := &node{next: nil, value: item}
	if l.isEmpty() {
		l.first = n
		l.last = n
	} else {
		f := l.first
		l.first = n
		l.first.next = f
	}
	l.size++
}

func (l *linkList) AddLast(item int) {
	if l.isEmpty() {
		l.AddFirst(item)
	} else {
		last := l.last
		l.last = &node{next: nil, value: item}
		last.next = l.last
		l.size++
	}
}

func (l *linkList) DeleteFirst() {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}
	next := l.first.next
	l.first.next = nil
	l.first = next
	l.size--
}

func (l *linkList) DeleteLast() {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}
	n := l.first

	if l.first == l.last {
		l.first = nil
	} else {
		for {
			if n.next == l.last {
				break
			} else {
				n = n.next
			}
		}
		l.last = n
		l.last.next = nil
	}
	l.size--
}

func (l *linkList) Contains(item int) bool {
	n := l.first
	for {
		if n == nil {
			break
		}

		if n.value == item {
			return true
		} else {
			n = n.next
		}
	}
	return false
}

func (l *linkList) IndexOf(item int) int {
	index := 0
	n := l.first
	for {
		if n == nil {
			break
		}

		if n.value == item {
			return index
		} else {
			n = n.next
			index++
		}
	}
	return -1
}

func (l *linkList) Size() int {
	return l.size
}

func (l *linkList) ToArray() []int {
	index := 0
	arr := make([]int, l.size)
	fmt.Println(arr)
	if l.first == l.last {
		arr[index] = l.first.value
	} else {
		n := l.first
		for {
			if n == nil {
				break
			} else {
				arr[index] = n.value
				n = n.next
				index++
			}
		}
	}
	return arr
}

func (l *linkList) Print() {
	fmt.Print("[")
	n := l.first
	for {
		if n != nil {
			fmt.Print(n.value)
			if n.next != nil {
				fmt.Print(", ")
			}
			n = n.next
		} else {
			break
		}
	}
	fmt.Println("]")
}
