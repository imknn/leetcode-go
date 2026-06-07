package array

/* 219. 存在重复元素 II */
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for idx, num := range nums {
		if _, ok := m[num]; ok {
			m[num] = append(m[num], idx)
		} else {
			n := make([]int, 1)
			n[0] = idx
			m[num] = n
		}
	}
	// 检查重复元素
	for _, v := range m {
		length := len(v)
		if length >= 2 {
			for i := 0; i < length; i++ {
				for j := i + 1; j < length; j++ {
					if v[j]-v[i] <= k {
						return true
					}
				}
			}
		}
	}
	return false
}
