package dynamic_programming

/* 509. 斐波那契数 */
func fib(N int) int {
	m, n := 1, 1
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	for i := 2; i < N; i++ {
		m, n = n, m+n
	}
	return n
}

//func fib(N int) int {
//	nums := [...]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040}
//	return nums[N]
//}
