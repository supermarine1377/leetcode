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
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "3rd",
			args: args{
				digits: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3},
			},
			want: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 4},
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

func Test_convertToInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 123,
		},
		{
			name: "2nd",
			args: args{
				nums: []int{5, 1, 2},
			},
			want: 512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToInt(tt.args.nums); got != tt.want {
				t.Errorf("convertToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow10(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				num: 2,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow10(tt.args.num); got != tt.want {
				t.Errorf("pow10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToArr(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1st",
			args: args{
				num: 123,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToArr(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_digit(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				num: 123,
			},
			want: 3,
		},
		{
			name: "2nd",
			args: args{
				num: 1000,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit(tt.args.num); got != tt.want {
				t.Errorf("digit() = %v, want %v", got, tt.want)
			}
		})
	}
}
