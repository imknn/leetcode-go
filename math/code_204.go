package math

/* 204. 计数质数 */

func countPrimes(n int) int {
	nums := make([]bool, n)
	count := 0
	for i := 2; i*i < n; i++ {
		if !nums[i-1] {
			for j := i * i; j < n; j += i {
				nums[j-1] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !nums[i-1] {
			count++
		}
	}
	return count
}
