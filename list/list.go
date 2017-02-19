package list

import (
	"fmt"
)

type LinkedList struct {
	first   *node
	last    *node
	current *node
	size    int
}

type node struct {
	previous *node
	Value    interface{}
	next     *node
}

func (n node)String() string {
	return fmt.Sprintf("previous: %s, value: %s, next: %s", n.previous.Value, n.Value, n.next.Value)
}

func NewLinkedList() *LinkedList {
	return &LinkedList{size:0}
}

func (l *LinkedList) add(value interface{}) {
	n := &node{Value:value}
	if l.size == 0 {
		l.first = n
		l.last = l.first
	} else {
		n.previous = l.last
		l.last.next = n
		l.last = n
	}
	l.size++
}

func (l *LinkedList) reset()  {
	l.First()
}

func (l *LinkedList) find(value interface{}) (*node, int) {
	node := l.First()
	for i := 0; i < l.Size(); i++ {
		if node.Value == value {
			return node, i
		}
		node = l.Next()
	}
	l.reset()
	return nil, -1
}

func (l *LinkedList) AddFirst(value interface{}) {
	n := &node{Value:value, next:l.first}
	if l.first != nil {
		l.first.previous = n
		l.first = n
		l.size++
	} else {
		l.add(value)
	}
}

func (l *LinkedList) AddBefore(element, value interface{}) error {
	nd, index := l.find(element)
	if index == -1 {
		return fmt.Errorf("%v element not found in linked list", element)
	}

	if l.first != nd {
		n := &node{Value:value, next:nd, previous:nd.previous}
		nd.previous.next = n
		nd.previous = n
		l.size++
	} else {
		l.AddFirst(value)
	}
	return nil
}

func (l *LinkedList) AddAfter(element, value interface{}) error {
	nd, index := l.find(element)
	if index == -1 {
		return fmt.Errorf("%v element not found in linked list", element)
	}

	if l.last != nd {
		n := &node{Value:value, next:nd.next, previous:nd}
		nd.next.previous = n
		nd.next = n
		l.size++
	} else {
		l.AddLast(value)
	}
	return nil
}

func (l *LinkedList) AddLast(value interface{}) {
	l.add(value)
}

func (l *LinkedList) Get(index int) interface{} {
	if index >= l.Size() {
		// Todo panic because index is greater than size
	}
	node := l.First()
	for i := 0; i < index; i ++ {
		node = l.Next()
	}
	l.reset()
	return node.Value
}

func (l *LinkedList) Find(value interface{}) int {
	_, index := l.find(value)// element was not found
	return index
}

func (l *LinkedList) Next() *node {
	if l.current == nil {
		// Todo Add panic if current is nil
	} else {
		l.current = l.current.next
	}
	return l.current
}

func (l * LinkedList) Delete(value interface{}) error {
	node, index := l.find(value)
	if index == -1 {
		return fmt.Errorf("%v element not found in linked list", value)
	}

	if l.first == node {
		l.first = node.next
		node.next.previous = nil
	} else if l.last == node {
		l.last = node.previous
		node.previous.next = nil

	} else {
		node.previous.next = node.next
		node.next.previous = node.previous
	}
	node = nil

	return nil
}

func (l *LinkedList) Prev() *node {
	if l.current == nil {
		// Todo Add panic functionality if current is nil
	} else {
		l.current = l.current.previous
	}
	return l.current
}

func (l *LinkedList) First() *node {
	l.current = l.first
	return l.current
}

func (l *LinkedList) Last() *node {
	l.current = l.last
	return l.current
}

func (l *LinkedList) Size() int {
	return l.size
}
