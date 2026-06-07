package string

import "bytes"

/* 38. 报数 */
func countAndSay(n int) string {
	var result = "1"
	var buffer bytes.Buffer
	for ; n > 1; n-- {
		length := len(result)
		var num = 1
		for i := 0; i < length; i++ {
			if i == length-1 || result[i] != result[i+1] {
				// 开始打印
				buffer.WriteByte(byte(num + '0'))
				buffer.WriteByte(result[i])
				num = 1
			} else {
				num++
			}
		}
		result = buffer.String()
		buffer.Reset()
	}
	return result
}
