package valid_perfect_square

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1st",
			args: args{
				num: 1,
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				num: 4,
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				num: 9,
			},
			want: true,
		},
		{
			name: "4th",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "5th",
			args: args{
				num: 25,
			},
			want: true,
		},
		{
			name: "6th",
			args: args{
				num: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
