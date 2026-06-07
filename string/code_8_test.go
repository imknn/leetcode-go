package string

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "示例1: 正数",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "示例2: 带空格的负数",
			args: args{s: "   -42"},
			want: -42,
		},
		{
			name: "示例3: 数字后有字母",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "示例4: 字母开头",
			args: args{s: "words and 987"},
			want: 0,
		},
		{
			name: "示例5: 负数溢出",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "空字符串",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "只有空格",
			args: args{s: "   "},
			want: 0,
		},
		{
			name: "正数溢出",
			args: args{s: "91283472332"},
			want: 2147483647,
		},
		{
			name: "前导零",
			args: args{s: "00000-42"},
			want: 0,
		},
		{
			name: "加号开头",
			args: args{s: "+1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
