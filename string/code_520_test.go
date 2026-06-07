package string

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{"1", "USA", true},
		{"2", "FlaG", false},
		{"3", "leetcode", true},
		{"4", "Google", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
