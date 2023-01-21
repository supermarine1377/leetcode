package best_time_to_buy_and_sell_a_stock

// Greedy solution
// Time Complexity: O(N^2)
// Space Compleity: O(1)
func maxProfit(prices []int) int {
	var result int
	length := len(prices)
	for length > 0 {
		sp := prices[length-1]
		prices = prices[0 : length-1]
		length--
		for _, bp := range prices {
			if result < sp-bp {
				result = sp - bp
			}
		}
	}
	return result
}
