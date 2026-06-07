package string

// 6. Z 字形变换
// 模拟 Z 字形排列，将字符分配到各行，最后拼接

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]byte, numRows)
	curRow, goingDown := 0, false

	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		// 到达顶部或底部时改变方向
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}
	return string(result)
}
