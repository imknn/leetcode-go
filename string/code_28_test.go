package string

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello", "ll"}, 2},
		{"2", args{"aaaaa", "bba"}, -1},
		{"3", args{"aaaaa", ""}, 0},
		{"4", args{"abcdabcdefg", "abcde"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
