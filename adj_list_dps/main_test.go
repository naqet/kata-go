package adjlistdps_test

import (
	"reflect"
	"testing"
)

type Edge struct {
	to     int
	weight int
}

func walk(list [][]Edge, curr int, target int, seen *[]bool, path *[]int) bool {
	if curr == target {
		*path = append(*path, curr)
		return true
	}

	if (*seen)[curr] {
		return false
	}

	if curr >= len(list) {
		return false
	}

	(*seen)[curr] = true
	*path = append(*path, curr)

	for _, node := range list[curr] {
		found := walk(list, node.to, target, seen, path)
		if found {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func dps(list [][]Edge, source int, target int) []int {
	path := []int{}
	seen := make([]bool, len(list))

	walk(list, source, target, &seen, &path)

	return path
}

func TestDps(t *testing.T) {
	list := [][]Edge{
		{
			{1, 1},
			{2, 3},
			{3, 5},
		},
		{
			{0, 1},
		},
		{
			{3, 4},
		},
		{
			{4, 5},
		},
		{},
	}

	path := dps(list, 0, 4)

	if !reflect.DeepEqual(path, []int{0, 2, 3, 4}) {
		t.Error(path)
		return
	}

	path = dps(list, 0, 10)

	if !reflect.DeepEqual(path, []int{}) {
		t.Error(path)
		return
	}
}
