package bit_manipulation

import (
	"reflect"
	"testing"
)

func Test_findNumOfValidWords(t *testing.T) {
	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				words:   []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
				puzzles: []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			},
			want: []int{1, 1, 3, 2, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
