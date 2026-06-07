package array

import "sort"

// 15. 三数之和
// 排序 + 双指针，去重
// 时间 O(n^2)，空间 O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 跳过重复
		}
		if nums[i] > 0 {
			break // 剪枝：最小的数已经大于 0
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // 跳过重复
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // 跳过重复
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
