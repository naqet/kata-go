package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func bfs(head Node, target int) bool {
	result := false
	q := Queue{[]*Node{}}

	q.enque(&head)

	for len(q.arr) > 0 {
        head := q.deque()

        if head == nil {
            continue;
        }

		if head.value == target {
			result = true
			break
		}
		if head.left != nil {
			q.enque(head.left)
		}

		if head.right != nil {
			q.enque(head.right)
		}
	}

	return result
}

func main() {
	node := Node{0, &Node{1, &Node{3, nil, nil}, &Node{4, nil, nil}}, &Node{2, &Node{5, nil, nil}, &Node{6, nil, nil}}}

	fmt.Println(bfs(node, 4));
	fmt.Println(bfs(node, 100));
}

type Queue struct {
	arr []*Node
}

func (q *Queue) enque(val *Node) {
	q.arr = append(q.arr, val)
}

func (q *Queue) deque() *Node {
	if len(q.arr) > 0 {
        node := q.arr[0]
		q.arr = q.arr[1:]
        return node;
	}

    return nil;
}
