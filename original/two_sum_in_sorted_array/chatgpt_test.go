package two_sum_in_sorted_array

import "testing"

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				target:  9,
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				target:  7,
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				numbers: []int{1, 2, 3, 4, 5},
				target:  10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); got != tt.want {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
