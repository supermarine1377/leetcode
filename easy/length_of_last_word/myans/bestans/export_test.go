package bestans

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
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
		{
			name: "3rd",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "4th",
			args: args{
				s: "abc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
