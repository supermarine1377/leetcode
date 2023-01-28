package remove_duplicates_from_sorted_array

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_removeDuplicates(t *testing.T) {
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
				nums: nil,
			},
			want: nil,
		},
		{
			name: "2nd",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "3rd",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "4th",
			args: args{
				nums: []int{1, 2, 2, 2, 2, 3, 4, 4, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "5th",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected diff: %s", diff)
			}
		})
	}
}
