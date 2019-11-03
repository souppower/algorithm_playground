package fibonacci

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

var memo = make(map[int]int, 30)

func fibonacci2(N int) int {
	if N < 2 {
		return N
	}

	if m, ok := memo[N]; ok {
		return m
	}

	memo[N] = fibonacci2(N-2) + fibonacci2(N-1)

	return memo[N]
}
