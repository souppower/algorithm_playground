package knapsack

func knapsack(n int, w []int, v []int, C int) int {
	if n == 0 || C == 0 {
		return 0
	} else if C < w[n-1] {
		return knapsack(n-1, w, v, C)
	}

	return max(
		v[n-1]+knapsack(n-1, w, v, C-w[n-1]),
		knapsack(n-1, w, v, C),
	)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
