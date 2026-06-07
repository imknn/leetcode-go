package string

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{{
		name: "test_58",
		args: "Hello World",
		want: 5,
	}, {
		name: "test_58_2",
		args: " Helloww ",
		want: 7,
	}, {
		name: "test_58_3",
		args: "Today is a nice day",
		want: 3,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
