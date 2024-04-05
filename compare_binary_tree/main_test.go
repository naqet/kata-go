package comparebinarytree_test

import "testing"

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func compare[T int](first *Node[T], second *Node[T]) bool {
    if first == nil && second == nil {
        return true
    }

    if first == nil || second == nil {
        return false
    }

    if first.value != second.value {
        return false
    }

    return compare(first.left, second.left) && compare(first.right, second.right);
}

func TestCompare(t *testing.T) {
	first := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 4}}, right: &Node[int]{value: 3}}
	second := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 4}}, right: &Node[int]{value: 3}}

    result := compare(&first, &second)

    if !result {
        t.Error()
        return;
    }

    third := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 5}}, right: &Node[int]{value: 3}}

    result = compare(&first, &third)

    if result {
        t.Error()
        return;
    }

}
