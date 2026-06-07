package array

/*349. 两个数组的交集*/
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, i := range nums1 {
		m[i] = true
	}
	n := make(map[int]bool)
	for _, i := range nums2 {
		if _, ok := m[i]; ok {
			n[i] = true
		}
	}
	result, idx := make([]int, len(n)), 0
	for k, _ := range n {
		result[idx] = k
		idx++
	}
	return result
}
