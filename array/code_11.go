package array

// 11. 盛最多水的容器
// 双指针：每次移动较短的一侧
// 时间 O(n)，空间 O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
