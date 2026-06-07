package math

/* 9. 回文数 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var num []int
	for x > 0 {
		num = append(num, x%10)
		x /= 10
	}

	for l, i := len(num), 0; i < l/2; i++ {
		if num[i] != num[l-1-i] {
			return false
		}
	}
	return true
}
