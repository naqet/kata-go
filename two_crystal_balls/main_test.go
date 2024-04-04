package twocrystalballs_test

import (
	"math"
	"testing"
)

func logic(arr []bool) int {
	var result = -1

    length := float64(len(arr));
    jump := math.Sqrt(length);

    point := jump;

    for ;point < length; point += jump {
        i := int(math.Floor(point))
        if arr[i] {
            break;
        }
	}

    lo := int(math.Floor(point - jump))
    start := 0.0;

    for lo < len(arr) && start <= jump {
        if arr[lo] {
            result = lo;
            break;
        }

        lo++;
        start++;
    }

	return result
}

type Test struct {
	arr      []bool
	expected int
}

var tests = []Test{
	{[]bool{false, false, false, true, true}, 3},
	{[]bool{false, true, true, true, true}, 1},
	{[]bool{true}, 0},
	{[]bool{false}, -1},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if result := logic(test.arr); result != test.expected {
			t.Errorf("Got: %d, Expected: %d", result, test.expected)
			break
		}
	}
}
