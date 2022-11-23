package myans

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1st",
			args: args{
				[]string{"", ""},
			},
			want: "",
		},
		{
			name: "2nd",
			args: args{
				[]string{"", "", ""},
			},
			want: "",
		},
		{
			name: "3rd",
			args: args{
				[]string{"common", "coupon", "count"},
			},
			want: "co",
		},
		{
			name: "4th",
			args: args{
				[]string{"dog", "domain", "dictionary"},
			},
			want: "d",
		},
		{
			name: "5th",
			args: args{
				[]string{"height", "angry"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
