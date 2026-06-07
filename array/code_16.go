package array

import (
	"math"
	"sort"
)

// 16. 最接近的三数之和
// 排序 + 双指针
// 时间 O(n^2)，空间 O(1)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 确保 math 被使用（抑制编译器警告）
var _ = math.MaxInt32
