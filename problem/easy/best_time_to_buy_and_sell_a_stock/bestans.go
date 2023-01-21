package best_time_to_buy_and_sell_a_stock

func maxProfit_SlidingWindow(prices []int) int {
	var result int
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > result {
				result = profit
			}
		} else {
			l = r
		}
		r++
	}
	return result
}
