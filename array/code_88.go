package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	x, y := 0, 0
	for ; x < m && y < n; x++ {
		if nums1[x] > nums2[y] {
			for b := m; b > x; b-- {
				nums1[b] = nums1[b-1]
			}
			nums1[x] = nums2[y]
			m++
			y++
		}
	}
	for y < n {
		nums1[x] = nums2[y]
		x++
		y++
	}
}
