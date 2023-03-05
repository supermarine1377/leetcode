package myans

import "testing"

type args struct {
	s string
}

type testcase struct {
	name string
	args args
	want int
}

func getTeseCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "2nd",
			args: args{
				s: "Hi Hi Youtube  ",
			},
			want: 7,
		},
	}
}

func Test_lengthOfLastWord(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLastWord_Arr(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord_Arr(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
