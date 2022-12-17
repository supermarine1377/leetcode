package bestans

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{1, 2, 4, 6},
				target: 2,
			},
			want: 1,
		},
		{
			args: args{
				nums:   []int{1, 2, 4, 6},
				target: 6,
			},
			want: 3,
		},
		{
			args: args{
				nums:   []int{1, 2, 4, 6},
				target: 3,
			},
			want: 2,
		},
		{
			args: args{
				nums:   []int{1, 2, 4, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
