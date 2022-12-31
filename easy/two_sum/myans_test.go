package twosum

import (
	"reflect"
	"sort"
	"testing"
)

type args struct {
	nums   []int
	target int
}

type testcase struct {
	name string
	args args
	want []int
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				nums:   []int{1, 2, 5, 3},
				target: 7,
			},
			want: []int{1, 2},
		},
		{
			name: "2nd",
			args: args{
				nums:   []int{3, 3, 0, 1},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "3rd",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
	}
}

func Test_twoSum(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			// Orders doesn't matter
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error: want %v; but got %v", tt.want, got)
			}
		})
	}
}

func Test_twoSum_slow(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum_slow(tt.args.nums, tt.args.target)
			// Orders doesn't matter
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error: want %v; but got %v", tt.want, got)
			}
		})
	}
}
