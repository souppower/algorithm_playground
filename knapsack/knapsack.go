package knapsack

func KP(n int, w []int, v []int, C int) int {
	memo := make([][]int, n+1)
	for m1 := range memo {
		memo[m1] = make([]int, C+1)
		for m2 := range memo[m1] {
			memo[m1][m2] = -1
		}
	}

	return knapsack(n, w, v, C, memo)
}

func knapsack(n int, w []int, v []int, C int, memo [][]int) int {
	if memo[n][C] != -1 {
		return memo[n][C]
	}

	if n == 0 || C == 0 {
		memo[n][C] =  0
	} else if C < w[n-1] {
		memo[n][C] =  knapsack(n-1, w, v, C, memo)
	} else {
		memo[n][C] =  max(
			v[n-1]+knapsack(n-1, w, v, C-w[n-1], memo),
			knapsack(n-1, w, v, C, memo),
		)
	}

	return memo[n][C]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
