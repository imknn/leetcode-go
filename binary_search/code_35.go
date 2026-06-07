package binary_search

/* 35. 搜索插入位置 */
func searchInsert(nums []int, target int) int {
	length := len(nums)
	i := 0
	for i < length && nums[i] < target {
		i++
	}
	return i
}
