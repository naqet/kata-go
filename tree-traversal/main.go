package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func recurse(node *Node, path *[]int) []int {
	if node == nil {
		return *path
	}

	recurse(node.left, path)

	*path = append(*path, node.value)

	recurse(node.right, path)
	return *path
}

func main() {
	node := Node{0, &Node{1, &Node{3, nil, nil}, &Node{4, nil, nil}}, &Node{2, &Node{5, nil, nil}, &Node{6, nil, nil}}}

	fmt.Println(recurse(&node, &[]int{}))
}
