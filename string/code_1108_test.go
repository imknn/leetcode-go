package string

import "testing"

func Test_defangIPaddr(t *testing.T) {
	tests := []struct {
		address string
		want    string
	}{
		{
			want:    "1[.]1[.]1[.]1",
			address: "1.1.1.1",
		},
		{
			want:    "255[.]100[.]50[.]0",
			address: "255.100.50.0",
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := defangIPaddr(tt.address); got != tt.want {
				t.Errorf("defangIPaddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
