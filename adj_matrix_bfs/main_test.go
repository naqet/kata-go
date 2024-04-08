package adjmatrixbfs_test

import (
	"reflect"
	"testing"
)

func bfs(list [][]int, source int, target int) []int {
    if len(list) <= target {
        return []int{}
    }
	seen := make([]bool, len(list))
	prev := make([]int, len(list))

	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true

	q := []int{source}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr == target {
			break
		}

		for node, value := range list[curr] {
			if value == 0 || seen[node] {
				continue
			}

			seen[node] = true
			prev[node] = curr
			q = append(q, node)
		}
	}

	curr := target
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) != 0 {
		length := len(out)

		for i := 0; i < length/2; i++ {
			out[i], out[length-1-i] = out[length-1-i], out[i]
		};

        out = append([]int{curr}, out...);
	}

	return out
}

func TestBts(t *testing.T) {
	matrix := [][]int{
		{0, 1, 4, 5, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 2, 0},
		{0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0},
	}

	out := bfs(matrix, 0, 4)

	if !reflect.DeepEqual(out, []int{0, 3, 4}) {
		t.Error(out)
		return
	}

	out = bfs(matrix, 0, 10)

	if !reflect.DeepEqual(out, []int{}) {
		t.Error(out)
		return
	}
}
