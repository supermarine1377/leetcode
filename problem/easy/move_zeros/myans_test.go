package move_zeros

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				nums: []int{0, 3, 1, 2, 0, 1, 2, 0},
			},
			want: []int{3, 1, 2, 1, 2, 0, 0, 0},
		},
		{
			name: "2nd",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			moveZeroes(nums)
			if diff := cmp.Diff(nums, tt.want); diff != "" {
				t.Errorf("unexpeced diff: %s", diff)
			}
		})
	}
}
