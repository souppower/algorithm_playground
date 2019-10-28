package quickSort

type List struct {
	Values []int
}

func (list *List) Sort() *List {
	sorted := quickSort(list.Values, 0, len(list.Values)-1)
	return &List{Values: sorted}
}

func quickSort(list []int, low int, high int) []int {
	if low > high {
		return list
	}

	p := partition(list, low, high)
	quickSort(list, low, p-1)
	quickSort(list, p+1, high)

	return list
}

func partition(list []int, low, high int) int {
	pivot := list[high]

	for i := low; i < high; i++ {
		if list[i] < pivot {
			list[i], list[low] = list[low], list[i]
			low++
		}
	}

	list[low], list[high] = list[high], list[low]

	return low
}
