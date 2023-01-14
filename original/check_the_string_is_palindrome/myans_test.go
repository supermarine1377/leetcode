package check_the_string_is_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				str: "geeksforgeeks",
			},
			want: false,
		},
		{
			name: "2nd",
			args: args{
				str: "abca",
			},
			want: false,
		},
		{
			name: "3rd",
			args: args{
				str: "madam",
			},
			want: true,
		},
		{
			name: "4th",
			args: args{
				str: "abba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
