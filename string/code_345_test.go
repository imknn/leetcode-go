package string

import "testing"

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"1", "hello", "holle"},
		{"2", "leetcode", "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
