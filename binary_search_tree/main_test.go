package binarysearchtree_test;

import "testing"

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func find[T int](node *Node[T], target T) bool {
    if node == nil {
        return false;
    }

    if node.value == target {
        return true;
    }

    if node.value < target && node.right != nil {
        return find(node.right, target)
    }

    if node.value > target && node.left != nil {
        return find(node.left, target)
    }

    return false;
}

func TestFind(t *testing.T) {
	tree := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 3}}, right: &Node[int]{value:4}}

    result := find(&tree, 4);

    if !result {
        t.Error();
        return;
    }

    result = find(&tree, 10);

    if result {
        t.Error();
        return;
    }
}
