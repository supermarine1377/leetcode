package valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				s: "abba",
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				s: "AbBa",
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "4th",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "5th",
			args: args{
				s: "  ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
