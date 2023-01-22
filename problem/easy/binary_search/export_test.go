package binary_search

import "testing"

type args struct {
	nums   []int
	target int
}

type testcase struct {
	name string
	args args
	want int
}

func getTeseCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				nums:   []int{-1, 2, 3, 5, 6, 7},
				target: 5,
			},
			want: 3,
		},
		{
			name: "2nd",
			args: args{
				nums:   []int{-1, 2, 3, 5, 6, 7},
				target: 2,
			},
			want: 1,
		},
		{
			name: "3rd",
			args: args{
				nums:   []int{-1, 2, 3, 5, 6, 7, 9},
				target: 7,
			},
			want: 5,
		},
		{
			name: "3rd",
			args: args{
				nums:   []int{-1, 2, 3, 5, 6, 7, 9},
				target: 1,
			},
			want: -1,
		},
	}
}

func Test_serach(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := serach(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("serach() = %v, want %v", got, tt.want)
			}
		})
	}
}
