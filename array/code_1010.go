package array

/* 1010. 总持续时间可被 60 整除的歌曲 */
func numPairsDivisibleBy60(time []int) int {
	dir := make([]int, 60)
	for _, t := range time {
		dir[t%60]++
	}
	count := 0
	for idx, t := range dir {
		if t > 0 {
			if idx == 0 || idx == 30 {
				count += (t * (t - 1)) / 2
			} else {
				count += t * dir[60-idx]
			}
		}
		if idx >= 30 {
			break
		}
	}
	return count
}
