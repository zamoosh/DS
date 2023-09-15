package linkList

import "fmt"

type node struct {
	next  *node
	value int
}

type linkList struct {
	first *node
	last  *node
}

func LinkList() linkList {
	l := linkList{first: nil, last: nil}
	return l
}

func (l *linkList) AddFirst(item int) {
	if l.first == nil {
		l.first = &node{next: nil, value: item}
		l.last = l.first
	} else {
		f := l.first
		l.first = &node{next: f, value: item}
	}
}

func (l *linkList) AddLast(item int) {
	if l.first == nil && l.last == nil {
		l.AddFirst(item)
	} else {
		n := &node{next: nil, value: item}
		l.last.next = n
		l.last = n
	}
}

func (l *linkList) DeleteFirst() {
	next := l.first.next
	l.first.next = nil
	l.first = next
}

func (l *linkList) DeleteLast() {
	n := l.first
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
