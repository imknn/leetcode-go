package hash

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args args
		want [][]string
	}{
		{args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
	}
	for _, tt := range tests {
		t.Run("49", func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
