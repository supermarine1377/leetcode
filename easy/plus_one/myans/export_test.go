package myans

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "2nd",
			args: args{
				digits: []int{1, 2, 9},
			},
			want: []int{1, 3, 0},
		},
		{
			name: "3rd",
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "4th",
			args: args{
				digits: []int{1, 9, 9},
			},
			want: []int{2, 0, 0},
		},
		{
			name: "5th",
			args: args{
				digits: []int{1},
			},
			want: []int{2},
		},
		{
			name: "6th",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
