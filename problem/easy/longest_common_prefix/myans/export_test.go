package myans

import "testing"

type args struct {
	strs []string
}

type testcase struct {
	name string
	args args
	want string
}

func getTestCases() []testcase {
	return []testcase{
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
}

func Test_longestCommonPrefix(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefixO2(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixO2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
