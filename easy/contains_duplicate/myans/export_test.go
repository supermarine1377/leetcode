package myans

import (
	"testing"
)

type args struct {
	nums []int
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
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "2nd",
			args: args{
				nums: []int{2, 1, 4, 3},
			},
			want: false,
		},
		{
			name: "3rd",
			args: args{
				nums: []int{4, 3, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "4th",
			args: args{
				nums: []int{3, 0, -2, -1, 2, 1},
			},
			want: false,
		},
		{
			name: "4th",
			args: args{
				nums: []int{3, 0, -2, -1, -2, 1},
			},
			want: true,
		},
	}
}

func Test_containsDuplicate(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicateO1(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicateO1(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicateO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
