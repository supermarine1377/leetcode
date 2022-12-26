package twosum

import (
	"reflect"
	"sort"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
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
	for _, tt := range tests {
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
