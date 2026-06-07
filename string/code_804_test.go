package string

import "testing"

func Test_uniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{"1", []string{"gin", "zen", "gig", "msg"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
