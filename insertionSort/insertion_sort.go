package insertionSort

type List struct {
	Values []int
}

func (l *List) Sort() {
	for ins := 1; ins < len(l.Values); ins++ {
		temp := l.Values[ins]
		comp := ins
		for comp > 0 && l.Values[comp-1] > temp {
			l.Values[comp] = l.Values[comp-1]
			comp--
		}
		l.Values[comp] = temp
	}
}
