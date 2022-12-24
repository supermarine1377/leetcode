package myans

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "2nd",
			args: args{
				nums: []int{2, 1, 4, 3},
			},
			want: false,
		},
		{
			name: "3rd",
			args: args{
				nums: []int{4, 3, 4, 2, 1},
			},
			want: true,
		},
		{
			name: "4th",
			args: args{
				nums: []int{3, 0, -2, -1, 2, 1},
			},
			want: false,
		},
		{
			name: "4th",
			args: args{
				nums: []int{3, 0, -2, -1, -2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
