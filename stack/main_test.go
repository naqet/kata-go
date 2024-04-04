package stack_test

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
    length int
    head *Node[T]
}

func (s *Stack[T]) push(val T) {
    s.length++;
    node := Node[T]{value: val}
    if s.head == nil {
        s.head = &node;
        return
    }

    node.prev = s.head;
    s.head = &node;
}

func (s *Stack[T]) pop() (T, error) {
    if s.head == nil {
        var zeroVal T;
        return zeroVal, errors.New("Stack is empty");
    }

    s.length--;

    popped := s.head;
    if s.length != 0 {
        s.head = s.head.prev;
    } else {
        s.head = nil;
    }
    
    return popped.value, nil;
}
