package array

/* 665. 非递减数列 */
func checkPossibility(nums []int) bool {
	length := len(nums)
	for j := 1; j < length; j++ {
		if nums[j-1] > nums[j] {
			l := nums[j-1]
			r := nums[j]
			// 统一设置成右值（小值）
			nums[j-1] = r
			nums[j] = r
			if !isAsc(nums, length) {
				// 统一设置成左值 （大值）
				nums[j-1] = l
				nums[j] = l
				return isAsc(nums, length)
			}
		}
	}
	return true
}

// 判断数列是否递增
func isAsc(nums []int, length int) bool {
	for i := 1; i < length; i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}
