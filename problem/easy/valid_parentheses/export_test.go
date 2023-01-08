/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start

package valid_parentheses

import "testing"

type args struct {
	s string
}
type testcase struct {
	name string
	args args
	want bool
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "4th",
			args: args{
				s: "([)]",
			},
			want: false,
		},
		{
			name: "5th",
			args: args{
				s: "(())",
			},
			want: true,
		},
		{
			name: "6th",
			args: args{
				s: "({})",
			},
			want: true,
		},
	}
}

func Test_isValid(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid_Myans(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid_Myans(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
