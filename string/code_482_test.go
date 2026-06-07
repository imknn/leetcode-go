package string

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"5F3Z-2e-9-w", 4}, "5F3Z-2E9W"},
		{"2", args{"2-5g-3-J", 2}, "2-5G-3J"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
