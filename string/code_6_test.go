package string

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{"示例1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"示例2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"示例3", "A", 1, "A"},
		{"单行", "AB", 1, "AB"},
		{"两行", "ABCD", 2, "ACBD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
