package string

import (
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"123", "456"}, "56088"},
		{"2", args{"2", "3"}, "6"},
		{"3", args{"0", "0"}, "0"},
		{"4", args{"0", "1000"}, "0"},
		{"5", args{"88888", "999"}, "88799112"},
		{"6", args{"123456789", "987654321"}, "121932631112635269"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
