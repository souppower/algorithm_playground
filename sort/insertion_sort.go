package main

func insertion_sort(list []int) []int {
	for ins := 1; ins < len(list); ins++ {
		temp := list[ins]
		comp := ins
		for comp > 0 && list[comp-1] > temp {
			list[comp] = list[comp-1]
			comp--
		}
		list[comp] = temp
	}

	return list
}
