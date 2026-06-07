package array

/**
 * 找到0就和右侧的非零数字交换
 */
func moveZeroes(nums []int) {
	length := len(nums)
	for x := 0; x < length; x++ {
		if nums[x] == 0 {
			y := x + 1
			for y < length && nums[y] == 0 {
				y++
			}
			if y < length {
				nums[x], nums[y] = nums[y], nums[x]
			} else {
				// 0 放到末尾，长度-1
				length--
			}
		}
	}
}
