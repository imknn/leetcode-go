package hash

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"1", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"2", "AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"3", "AAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}

}
