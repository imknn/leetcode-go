package math

/* 268. 缺失数字 */
func missingNumber(nums []int) int {
	sum := len(nums)
	for i, v := range nums {
		sum += i - v
	}
	return sum
}
