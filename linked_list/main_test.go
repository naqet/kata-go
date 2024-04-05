package linkedlist_test

import (
	"errors"
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

func (l *LinkedList[T]) getAt(idx int) (T, error) {
    if idx >= l.length || l.length == 0 {
        var zeroVal T;
        return zeroVal, errors.New("Idx greater than length");
    }

    curr := l.head;

    for i := 1; i <= idx; i++ {
        curr = curr.next;
    }


    return curr.value, nil;
}

func (l *LinkedList[T]) insertAt(val T, idx int) error {
    if idx > l.length {
        return errors.New("Idx greater than length");
    } else if idx == l.length {
        l.append(val)
    } else if idx == 0 {
        l.prepend(val)
    }

    node := Node[T]{value: val};

    curr := l.head;
    for i := 1; i <= idx; i++ {
        curr = curr.next;
    }

    if curr.prev != nil {
        curr.prev.next = &node;
        node.prev = curr.prev
    }
    curr.prev = &node;
    node.next = curr;

    return nil
}

func TestInserAt(t *testing.T) {
	l := LinkedList[int]{}
	l.append(1)
	l.append(2)
	l.append(4)

    err := l.insertAt(3, 2)

    if err != nil {
        t.Error();
        return
    }

    val, err := l.getAt(2)

    if err != nil || val != 3 {
        t.Error();
        return
    }
}

func TestGetAt(t *testing.T) {
	l := LinkedList[int]{}
	l.append(1)
	l.append(2)
	l.append(3)

    val, err := l.getAt(0)

	if err != nil || val != 1 {
		return
	}

    val, err = l.getAt(2)

	if err != nil || val != 3 {
		return
	}
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
