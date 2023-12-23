package main

import "fmt"

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func quicksort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	idx := partition(arr, lo, hi)

	quicksort(arr, lo, idx-1)
	quicksort(arr, idx+1, hi)
}

func main() {
	arr := []int{3, 5, 2, 7, 6, 3}

	fmt.Println(arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
