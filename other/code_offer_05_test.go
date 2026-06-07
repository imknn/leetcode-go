package other

import "testing"

// 剑指 Offer 50. 第一个只出现一次的字符
func Test_replaceSpace(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"1", "We are happy.", "We%20are%20happy."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
