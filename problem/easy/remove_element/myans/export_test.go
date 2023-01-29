package myans

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type args struct {
	nums []int
	val  int
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
				nums: nil,
				val:  0,
			},
			want: nil,
		},
		{
			name: "2nd",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: []int{},
		},
		{
			name: "3rd",
			args: args{
				nums: []int{1},
				val:  0,
			},
			want: []int{1},
		},
		{
			name: "4th",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: []int{2, 2},
		},
	}
}

func Test_removeElement(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			// order dosen't matter
			sort.Ints(got)
			sort.Ints(tt.want)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}

func Test_removeElement_notUsingTwoPointer(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement_notUsingTwoPointers(tt.args.nums, tt.args.val)
			// order dosen't matter
			sort.Ints(got)
			sort.Ints(tt.want)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}
