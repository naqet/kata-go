package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func compare(a *Node, b *Node) bool {
    if a == nil && b == nil {
        return true;
    }

    if a == nil || b == nil {
        return false
    }

    if a.value != b.value {
        return false
    }

    return compare(a.left, b.left) && compare(a.right, b.right);
}

func main() {
	tree := Node{0, &Node{1, nil, nil}, &Node{2, nil, nil}}
	tree2 := Node{0, &Node{1, &Node{2, nil, nil}, nil}, nil}

    fmt.Println(compare(&tree, &tree2))

}
