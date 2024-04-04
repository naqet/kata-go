package bubblesort_test

import (
	"reflect"
	"testing"
)

func logic(arr []int){
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr) - 1 - i; j++ {
            if arr[j] > arr[j + 1] {
                tmp := arr[j + 1]
                arr[j + 1] = arr[j]
                arr[j] = tmp;
            }
        }
    }
}

type Test struct {
    arr []int
    expected []int
}

var tests = []Test{
    {[]int{3,2,1}, []int{1,2,3}},
    {[]int{1}, []int{1}},
    {[]int{6,5,4,2,1}, []int{1,2,4,5,6}},
}

func TestLogic(t *testing.T) {
    for _, test := range tests {
        logic(test.arr)
        if !reflect.DeepEqual(test.arr, test.expected) {
            t.Errorf("Got: %d, Expected: %d", test.arr, test.expected);
            break;
        }
    }
}
