package mazesolver_test

import (
	"reflect"
	"testing"
)

type Point struct {
	x int
	y int
}

var dirs = []Point{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func walk(p Point, maze [][]rune, path *[]Point, seen [6][6]bool) bool {
	if len(maze) <= p.y || len(maze[p.y]) <= p.x {
		return false
	}

	val := maze[p.y][p.x]

	if val == '#' {
		return false
	}

	if val == 'E' {
		*path = append(*path, p)
		return true
	}

	if seen[p.y][p.x] {
		return false
	}

	seen[p.y][p.x] = true
	*path = append(*path, p)

	for _, dir := range dirs {
		resolved := walk(Point{p.x + dir.x, p.y + dir.y}, maze, path, seen)

		if resolved {
			return true
		}
	}

	*path = (*path)[len(*path)-1:]

	return false
}

var test = [][]rune{
	{'#', '#', 'E', '#', 'S', '#'},
	{'#', '.', '.', '#', '.', '#'},
	{'#', '.', '#', '#', '.', '#'},
	{'#', '.', '#', '#', '.', '#'},
	{'#', '.', '.', '.', '.', '#'},
	{'#', '#', '#', '#', '#', '#'},
}

var expected = []Point{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {1, 2}, {1, 1}, {2, 1}, {2, 0}}

func TestLogic(t *testing.T) {
	path := []Point{}
	walk(Point{4, 0}, test, &path, [6][6]bool{})

	if !reflect.DeepEqual(path, expected) {
		t.Error("Not valid!", path, expected)
		return
	}
}
