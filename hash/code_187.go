package hash

/* 187. 重复的DNA序列 */
func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	length := len(s)
	if length < 10 {
		return result
	}
	m := make(map[int]int)
	// 二进制表示 A-00 C-01 G-10 和 T-11
	target := make([]int, 0)
	for i := 0; i < length-9; i++ {
		target = append(target, 0)
		for j := 0; j < 10; j++ {
			ch := s[i+j]
			switch ch {
			case 'A':
				target[i] = target[i]<<2 + 0
			case 'C':
				target[i] = target[i]<<2 + 1
			case 'G':
				target[i] = target[i]<<2 + 2
			case 'T':
				target[i] = target[i]<<2 + 3
			}
		}
		if m[target[i]] == 0 {
			m[target[i]] = 1
		} else if m[target[i]] == 1 {
			m[target[i]] = -length // 标记为已输出
			result = append(result, s[i:i+10])
		}
	}
	return result
}
