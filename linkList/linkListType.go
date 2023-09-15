package linkList

import (
	"fmt"
	"log"
)

type LinkList struct {
	first *node
	last  *node
	size  int
}

func NewLinkListType() LinkList {
	l := LinkList{first: nil, last: nil}
	return l
}

func (l *LinkList) isEmpty() bool {
	if l.first != nil {
		return false
	}
	return true
}

func (l *LinkList) AddFirst(item int) {
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

func (l *LinkList) AddLast(item int) {
	if l.isEmpty() {
		l.AddFirst(item)
	} else {
		last := l.last
		l.last = &node{next: nil, value: item}
		last.next = l.last
		l.size++
	}
}

func (l *LinkList) DeleteFirst() {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}
	next := l.first.next
	l.first.next = nil
	l.first = next
	l.size--
}

func (l *LinkList) DeleteLast() {
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

func (l *LinkList) Contains(item int) bool {
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

func (l *LinkList) IndexOf(item int) int {
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

func (l *LinkList) Size() int {
	return l.size
}

func (l *LinkList) ToArray() []int {
	index := 0
	arr := make([]int, l.size)
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

func (l *LinkList) Reverse() {
	items := l.ToArray()

	for {
		n := l.first
		if n == nil {
			break
		}
		l.DeleteFirst()
	}

	for i := 0; i < len(items); i++ {
		l.AddFirst(items[i])
	}
}

func (l *LinkList) KthNodeFromEndOld(k int) int {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}
	if k < 1 || k > l.size {
		log.Panic("k is invalid")
	}

	k--
	l.Reverse()
	arr := l.ToArray()
	return arr[k]
}

func (l *LinkList) KthNodeFromEnd(k int) int {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}

	a := l.first
	b := l.first
	for i := 0; i < k-1; i++ {
		b = b.next
		if b == nil {
			log.Panic("k is invalid")
		}
	}

	for {
		if b == l.last {
			break
		}

		a = a.next
		b = b.next
	}

	return a.value
}

func (l *LinkList) PrintMiddleOld() {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}

	arr := l.ToArray()
	length := len(arr)
	if length%2 == 0 {
		fmt.Printf("[%v, %v]\n", arr[length/2-1], arr[length/2])
	} else {
		fmt.Printf("[%v]\n", arr[length/2])
	}
}

func (l *LinkList) PrintMiddle() {
	if l.isEmpty() {
		log.Panic("linklist is empty")
	}

	count := 0
	a := l.first
	b := l.first

	for {
		if b != l.last && b.next != l.last {
			b = b.next.next
			a = a.next
			count += 2
		} else {
			break
		}
	}

	if b == l.last {
		fmt.Printf("[%v]\n", a.value)
	} else {
		fmt.Printf("[%v, %v]\n", a.value, a.next.value)
	}
}

func (l *LinkList) Print() {
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
