package main

import "fmt"

func search(arr []int, target int) int {
	idx := -1

	lo, hi := 0, len(arr)

	for lo < hi {
		mid := lo + ((hi - lo) / 2)
		val := arr[mid]

		if target == val {
			idx = mid
			break
		} else if target < val {
			hi = mid
		} else if target > val {
			lo = mid + 1
		}
	}

	return idx
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 6
	fmt.Println(search(arr, target))
	target = 9
	fmt.Println(search(arr, target))
}
