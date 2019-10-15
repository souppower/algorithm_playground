package main

func binary_search(list []int, searchValue int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == searchValue {
			return mid
		} else if list[mid] < searchValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
