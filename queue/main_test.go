package queue_test

import "errors"
type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
    length int
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) enqueue(val T) {
    q.length++;
	node := Node[T]{value: val}
	if q.tail == nil {
		q.tail = &node
		q.head = q.tail
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue[T]) deque() (T, error) {
	if q.head == nil {
		var zeroVal T
		return zeroVal, errors.New("Queue is empty")
	}
    q.length--;

    prevHead := q.head;
    q.head = q.head.next;

    prevHead.next = nil;

    if q.length == 0 {
        q.tail = nil;
    }

	return prevHead.value, nil
}

func (q *Queue[T]) peek() (T, error) {
	if q.head == nil {
		var zeroVal T
		return zeroVal, errors.New("Queue is empty")
	}

    return q.head.value, nil;
}
