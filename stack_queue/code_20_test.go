package stack_queue

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"1", "()", true},
		{"2", "()[]{}", true},
		{"3", "(]", false},
		{"4", "([)]", false},
		{"5", "{[]}", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
