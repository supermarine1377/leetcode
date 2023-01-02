package next_greater_element

import (
	"reflect"
	"testing"
)

type args struct {
	nums1 []int
	nums2 []int
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
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			name: "2nd",
			args: args{
				nums1: []int{2, 4},
				nums2: []int{1, 2, 3, 4},
			},
			want: []int{3, -1},
		},
	}
}

func Test_nextGreaterElement_O2(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement_O2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement_O2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElement_O_LenNum1_LenNum2(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement_O_LenNum1_LenNum2(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement_O2() = %v, want %v", got, tt.want)
			}
		})
	}
}
