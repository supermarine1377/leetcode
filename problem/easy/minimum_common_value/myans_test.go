package minimum_common_value

import "testing"

type args struct {
	nums1 []int
	nums2 []int
}

type testcase struct {
	name string
	args args
	want int
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "2nd",
			args: args{
				nums1: []int{2, 3, 4},
				nums2: []int{1, 3, 4},
			},
			want: 3,
		},
		{
			name: "2nd",
			args: args{
				nums1: []int{1, 3, 4},
				nums2: []int{2, 5, 6},
			},
			want: -1,
		},
	}
}

func Test_getCommon(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommon(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
