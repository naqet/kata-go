package quicksort_test

import (
	"reflect"
	"testing"
)

func partition(arr *[]int, lo int, hi int) int {
    pivot := (*arr)[hi];
    idx := lo - 1;
    for i := lo; i < hi; i++ {
        if (*arr)[i] <= pivot {
            idx++
            tmp := (*arr)[i]
            (*arr)[i] = (*arr)[idx]
            (*arr)[idx] = tmp;
        }
    }

    idx++
    (*arr)[hi] = (*arr)[idx]
    (*arr)[idx] = pivot;

    return idx;
}

func quicksort(arr *[]int, lo int, hi int) {
    if lo >= hi {
        return
    }

    pivot := partition(arr, lo, hi);

    quicksort(arr, lo, pivot - 1)
    quicksort(arr, pivot + 1, hi)
}

func TestQuickSort(t *testing.T) {
    unsorted := []int{5,3,4,2,1};
    expected := []int{1,2,3,4,5};
    quicksort(&unsorted, 0, len(unsorted) - 1);

    if !reflect.DeepEqual(unsorted, expected) {
        t.Errorf("Not equal: %x ||| %x", unsorted, expected);
        return
    }

    unsorted = []int{1};
    expected = []int{1};
    quicksort(&unsorted, 0, len(unsorted) - 1);

    if !reflect.DeepEqual(unsorted, expected) {
        t.Errorf("Not equal: %x ||| %x", unsorted, expected);
        return;
    }
}
