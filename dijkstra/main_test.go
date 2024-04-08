package dijkstra_test

import (
	"math"
	"reflect"
	"testing"
)

type Edge struct {
	to     int
	weight float64
}

func logic(list [][]Edge, source int, target int) []int {
	if target >= len(list) {
		return []int{}
	}

	seen := make([]bool, len(list))
	dists := make([]float64, len(list))
	prev := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		dists[i] = math.Inf(0)
		prev[i] = -1
	}

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)

		seen[curr] = true

		for _, edge := range list[curr] {
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.weight

			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	out := []int{}
	curr := target

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	length := len(out)

	if length != 0 {
		for i := 0; i < length/2; i++ {
			out[i], out[length-1-i] = out[length-1-i], out[i]
		}
		out = append([]int{curr}, out...)
	}

	return out
}

func hasUnvisited(seen []bool, dists []float64) bool {
	result := false
	for i := 0; i < len(seen) && i < len(dists); i++ {
		if !seen[i] {
			result = true
			break
		}
	}

	return result
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	lowest := math.Inf(0)

	for i := 0; i < len(seen) && i < len(dists); i++ {
		if !seen[i] && dists[i] < lowest {
			lowest = dists[i]
			idx = i
		}
	}

	return idx
}

func TestDijkstra(t *testing.T) {
	list := [][]Edge{
		{
			{1, 1},
			{2, 3},
			{3, 5},
		},
		{
			{0, 1},
			{4, 1},
		},
		{
			{3, 4},
		},
		{
			{4, 5},
		},
		{
			{3, 2},
		},
	}

	path := logic(list, 0, 4)

	if !reflect.DeepEqual(path, []int{0, 1, 4}) {
		t.Error(path)
		return
	}

	path = logic(list, 0, 4)

	if !reflect.DeepEqual(path, []int{0, 1, 4}) {
		t.Error(path)
		return
	}

	path = logic(list, 0, 3)

	if !reflect.DeepEqual(path, []int{0, 1, 4, 3}) {
		t.Error(path)
		return
	}
}
