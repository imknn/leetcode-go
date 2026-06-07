package math

import (
	"bytes"
	"izqy.top/common"
)

// 168. Excel表列名称
func convertToTitle(columnNumber int) string {
	var buffer bytes.Buffer
	for columnNumber > 0 {
		columnNumber--
		buffer.WriteByte(byte('A' + (columnNumber % 26)))
		columnNumber /= 26
	}

	common.Reverse(&buffer)
	return buffer.String()
}
