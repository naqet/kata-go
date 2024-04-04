package binarysearch_test

import "testing"

func logic(arr []int, target int) int {
	result := -1

	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := arr[mid]

		if val == target {
			result = mid
			break
		} else if val > target {
			hi = mid - 1
		} else if val < target {
			lo = mid + 1
		}
	}
	return result
}

type Test struct {
	arr      []int
	target   int
	expected int
}

var tests = []Test{
	{[]int{1, 2, 3, 4, 5}, 3, 2},
	{[]int{1, 2, 3, 4, 5}, 8, -1},
	{[]int{}, 8, -1},
	{[]int{1}, 1, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if result := logic(test.arr, test.target); result != test.expected {
			t.Errorf("Got: %d, Expected: %d", result, test.expected)
			break
		}
	}
}
