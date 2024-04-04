package linkedlist_test

import (
	"testing"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (l *LinkedList[T]) prepend(val T) {
	if l.length == 0 {
		node := &Node[T]{value: val}
		l.head = node
		l.tail = l.head
        l.length++
		return
	}

	node := &Node[T]{value: val, next: l.head}
	l.head.prev = node
	l.head = node
	l.length++
}

func (l *LinkedList[T]) append(val T) {

	if l.length == 0 {
		node := &Node[T]{value: val}
		l.head = node
		l.tail = l.head
        l.length++
		return
	}

	node := &Node[T]{value: val, prev: l.tail}
	l.tail.next = node
	l.tail = node
    l.length++
}

func TestPrepend(t *testing.T) {
	l := LinkedList[int]{}
	l.prepend(1)

	if l.head == nil || l.head.value != 1 || l.length != 1 || l.head != l.tail {
		t.Error()
		return
	}
	initial := l.head

	l.prepend(2)

	if l.head.value != 2 || l.head.next.value != 1 || l.head.next != initial || l.head.next.prev != l.head || l.length != 2 {
		t.Error()
		return
	}
}

func TestAppend(t *testing.T) {
	l := LinkedList[int]{}
	l.append(1)

	if l.head == nil || l.head.value != 1 || l.length != 1 || l.head != l.tail || l.tail == nil {
		t.Error()
		return
	}

	l.append(2)

	if l.tail.value != 2 || l.tail.prev.value != 1 || l.head.next != l.tail || l.tail.prev != l.head || l.length != 2 {
		t.Error()
		return
	}
}
