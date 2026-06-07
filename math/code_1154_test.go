package math

import "testing"

func Test_dayOfYear(t *testing.T) {
	tests := []struct {
		name string
		date string
		want int
	}{
		{"1", "2019-01-09", 9},
		{"2", "2019-02-10", 41},
		{"3", "2003-03-01", 60},
		{"4", "2004-03-01", 61},
		{"5", "2016-02-29", 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfYear(tt.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
