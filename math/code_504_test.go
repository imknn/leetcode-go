package math

import (
	"testing"
)

func Test_convertToBase7(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"1", 100, "202"},
		{"2", -7, "-10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7(tt.num); got != tt.want {
				t.Errorf("convertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
