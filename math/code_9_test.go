package math

import "testing"

func TestIsPalindrome(t *testing.T) {
	if isPalindrome(121) {
		t.Log("ok")
	}
	if isPalindrome(-121) {
		t.Fail()
	}
	if isPalindrome(10) {
		t.Fail()
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"1", 121, true},
		{"2", -121, false},
		{"3", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
