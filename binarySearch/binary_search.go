package binarySearch

type List struct {
	Values []int
}

func (l *List) Search(searchValue int) int {
	left := 0
	right := len(l.Values) - 1

	for left <= right {
		mid := (left + right) / 2

		if l.Values[mid] == searchValue {
			return mid
		} else if l.Values[mid] < searchValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
