package array

/* 832. 翻转图像 */
func flipAndInvertImage(A [][]int) [][]int {
	for _, m := range A {
		length := len(m)
		for i := 0; ; i++ {
			j := length - i - 1
			if i == j {
				if m[i] == 1 {
					m[i] = 0
				} else {
					m[i] = 1
				}
				break
			}
			if i > j {
				break
			}
			// 水平翻转
			m[i], m[j] = m[j], m[i]
			// 反转
			if m[i] == 1 {
				m[i] = 0
			} else {
				m[i] = 1
			}
			if m[j] == 1 {
				m[j] = 0
			} else {
				m[j] = 1
			}
		}
	}
	return A
}
