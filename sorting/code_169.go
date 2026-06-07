package sorting

/* 169. 求众数 */
func majorityElement(nums []int) int {
	m := make(map[int]int)
	s := len(nums) / 2
	for _, n := range nums {
		if _, ok := m[n]; ok {
			m[n]++
		} else {
			m[n] = 1
		}
		if m[n] > s {
			return n
		}
	}
	return nums[0]
}
