package hash

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j-i <= k {
				abs := nums[j] - nums[i]
				if abs < 0 {
					abs = -abs
				}
				if abs <= t {
					return true
				}
			}
		}
	}
	return false
}
