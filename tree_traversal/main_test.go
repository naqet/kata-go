package treetraversal_test

import (
	"reflect"
	"testing"
)

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func traverse(node *Node[int], path *[]int) {
    if node == nil {
        return
    }

    if node.left != nil {
        traverse(node.left, path)
    }

    (*path) = append((*path), node.value);

    if node.right != nil {
        traverse(node.right, path)
    }
}

func TestTraverse(t *testing.T) {
    node := Node[int]{value: 1, left: &Node[int]{value: 2}, right: &Node[int]{value: 3}}

    path := []int{}
    traverse(&node, &path);
    expected := []int{2,1,3}

    if !reflect.DeepEqual(path, expected) {
        t.Errorf("Got: %x, expected: %x", path, expected)
        return;
    }
}
