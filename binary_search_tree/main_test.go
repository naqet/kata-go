package binarysearchtree_test

import (
	"reflect"
	"testing"
)

type Node[T any] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

func find[T int](node *Node[T], target T) bool {
	if node == nil {
		return false
	}

	if node.value == target {
		return true
	}

	if node.value < target && node.right != nil {
		return find(node.right, target)
	}

	if node.value > target && node.left != nil {
		return find(node.left, target)
	}

	return false
}

func TestFind(t *testing.T) {
	tree := Node[int]{value: 1, left: &Node[int]{value: 2, left: &Node[int]{value: 3}}, right: &Node[int]{value: 4}}

	result := find(&tree, 4)

	if !result {
		t.Error()
		return
	}

	result = find(&tree, 10)

	if result {
		t.Error()
		return
	}
}

func insert[T int](tree *Node[T], val T) {
	if tree == nil {
		return
	}

	if val <= tree.value && tree.left != nil {
		insert(tree.left, val)
		return
	}
	if val > tree.value && tree.right != nil {
		insert(tree.right, val)
		return
	}

	if tree.left == nil && tree.right == nil {
		node := &Node[T]{value: val}
		if val <= tree.value {
			tree.left = node
		} else {
			tree.right = node
		}
	}
}

func TestInsert(t *testing.T) {
	tree := Node[int]{value: 10, left: &Node[int]{value: 5}, right: &Node[int]{value: 14}}
	expected := Node[int]{value: 10, left: &Node[int]{value: 5, right: &Node[int]{value: 8}}, right: &Node[int]{value: 14}}

	insert(&tree, 8)

	if !reflect.DeepEqual(tree, expected) {
		t.Error()
		return
	}

	tree = Node[int]{value: 10, left: &Node[int]{value: 5}, right: &Node[int]{value: 14}}
	expected = Node[int]{value: 10, left: &Node[int]{value: 5, right: &Node[int]{value: 8, right: &Node[int]{value: 9}}}, right: &Node[int]{value: 14, left: &Node[int]{value: 12}}}

	insert(&tree, 8)
	insert(&tree, 9)
	insert(&tree, 12)

	if !reflect.DeepEqual(tree, expected) {
		t.Error()
		return
	}
}

func delete[T int](tree **Node[T], val T) {
	if tree == nil {
		return
	}

	if val < (*tree).value {
		delete(&((*tree).left), val)
		return
	}

	if val > (*tree).value {
		delete(&((*tree).right), val)
		return
	}

	if (*tree).left == nil && (*tree).right == nil {
		*tree = nil
		return
	}

	if (*tree).left != nil && (*tree).right == nil {
		*tree = (*tree).left
		return
	}

	if (*tree).right != nil && (*tree).left == nil {
		*tree = (*tree).right
		return
	}

	parent := tree
	maxNode := &((*parent).left)
	for maxNode != nil && (*maxNode).right != nil {
		parent = maxNode
		maxNode = &((*parent).right)
	}

	if (*maxNode).left != nil {
		(*parent).right = (*maxNode).left
	}

	(*tree).value = (*maxNode).value
	*maxNode = nil
}

func traverse(node *Node[int], path *[]int) {
	if node == nil {
		return
	}

	if node.left != nil {
		traverse(node.left, path)
	}

	(*path) = append((*path), node.value)

	if node.right != nil {
		traverse(node.right, path)
	}
}

type Test struct {
	node     *Node[int]
	val      int
	expected []int
}

var tests = []Test{
	{
		&Node[int]{value: 10, left: &Node[int]{value: 5}, right: &Node[int]{value: 15}},
		5,
		[]int{10, 15},
	},
	{
		&Node[int]{value: 10, left: &Node[int]{value: 5, right: &Node[int]{value: 8}}, right: &Node[int]{value: 15}},
		5,
		[]int{8, 10, 15},
	},

	{
		&Node[int]{value: 10, left: &Node[int]{value: 5, left: &Node[int]{value: 2}, right: &Node[int]{value: 8, left: &Node[int]{value: 6}}}, right: &Node[int]{value: 15}},
		5,
		[]int{2, 6, 8, 10, 15},
	},
	{
		&Node[int]{value: 10, left: &Node[int]{value: 5, left: &Node[int]{value: 2, right: &Node[int]{value: 3}}, right: &Node[int]{value: 8, left: &Node[int]{value: 6}}}, right: &Node[int]{value: 15}},
		5,
		[]int{2, 3, 6, 8, 10, 15},
	},
	{
		&Node[int]{value: 10, left: &Node[int]{value: 5}, right: &Node[int]{value: 15, left: &Node[int]{value: 11, right: &Node[int]{value: 13}}, right: &Node[int]{value: 20}}},
		11,
		[]int{5,10,13,15,20},
	},
}

func TestDelete(t *testing.T) {
	for _, test := range tests {
		path := []int{}
		delete(&test.node, test.val)
		traverse(test.node, &path)

		if !reflect.DeepEqual(path, test.expected) {
			t.Error(path, test.expected)
			break
		}
	}
}
