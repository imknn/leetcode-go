package array

/* 26. 删除排序数组中的重复项 */
func removeDuplicates(nums []int) int {
	head, tail := 0, 1
	for l := len(nums); tail < l; tail++ {
		if nums[head] != nums[tail] {
			head++
			nums[head] = nums[tail]
		}
	}
	return head + 1
}
