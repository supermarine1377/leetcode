package find_is_subset_or_not

import "testing"

func Test_isSubset(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				arr1: []int{11, 1, 13, 21, 3, 7},
				arr2: []int{11, 3, 7, 1},
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				arr1: []int{1, 2, 3, 4, 5, 6},
				arr2: []int{1, 2, 4},
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				arr1: []int{10, 5, 2, 23, 19},
				arr2: []int{19, 5, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubset(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("isSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
