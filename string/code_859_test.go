package string

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"ab", "ba"}, true},
		{"2", args{"ab", "ab"}, false},
		{"3", args{"aaaaaaabc", "aaaaaaacb"}, true},
		{"4", args{"", "aa"}, false},
		{"5", args{"aaaaaaabccb", "aaaaaaacbbc"}, false},
		{"6", args{"aaaaa", "aaaaa"}, true},
		{"7", args{"abab", "abab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
