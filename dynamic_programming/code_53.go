package dynamic_programming

/* 53. 最大子序和 */

// 设置 f(n) 为前n个数的最大子串和，则 f(n) = max(f(n-1) + num[n], num[n])
// 这里不需要新的空间来保存 前n个数的最大子串和，使用已经读取过的nums

func maxSubArray(nums []int) int {
	max, length := nums[0], len(nums)
	for i := 1; i < length; i++ {
		if nums[i] < nums[i-1]+nums[i] { // 这里移项就是 nums[i-1] > 0
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
