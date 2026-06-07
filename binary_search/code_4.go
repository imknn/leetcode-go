package binary_search

import "math"

// 4. 寻找两个正序数组的中位数
// 二分查找：在较短数组上二分，找到正确的分割点
// 要求时间复杂度 O(log(m+n))

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left <= right {
		// partition1: nums1 左半部分的个数
		partition1 := (left + right) / 2
		// partition2: 总左半部分个数 - partition1
		partition2 := (m+n+1)/2 - partition1

		// 获取边界值，越界用 ±∞
		maxLeft1 := math.MinInt64
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}
		minRight1 := math.MaxInt64
		if partition1 < m {
			minRight1 = nums1[partition1]
		}
		maxLeft2 := math.MinInt64
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}
		minRight2 := math.MaxInt64
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			// 找到正确的分割点
			if (m+n)%2 == 1 {
				return float64(max(maxLeft1, maxLeft2))
			}
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
		} else if maxLeft1 > minRight2 {
			right = partition1 - 1
		} else {
			left = partition1 + 1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
