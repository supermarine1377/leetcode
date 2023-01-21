package best_time_to_buy_and_sell_a_stock

import "testing"

type args struct {
	prices []int
}

type testcase struct {
	name string
	args args
	want int
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "2nd",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "3rd",
			args: args{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
	}
}

func Test_maxProfit(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit_SlidingWindow(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_SlidingWindow(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
