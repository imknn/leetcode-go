package string

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{"1", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"2", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reverseString(tt.s); !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("fib() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}
