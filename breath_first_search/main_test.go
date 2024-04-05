package treetraversal_test

import (
	"testing"
)

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func bfs(node Node[int], target int) bool {
    var result bool;
	queue := []Node[int]{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
        if curr.value == target {
            result = true;
            break;
        }

		if curr.left != nil {
			queue = append(queue, *curr.left)
		}

		if curr.right != nil {
			queue = append(queue, *curr.right)
		}
	}

    return result;
}

func TestTraverse(t *testing.T) {
	node := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 4}}, right: &Node[int]{value: 3}}

    result := bfs(node, 3);

    if !result {
        t.Error()
        return;
    }

    result = bfs(node, 10);

    if result {
        t.Error()
        return;
    }
}
