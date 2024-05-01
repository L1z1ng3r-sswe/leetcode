func maxProfit(prices []int) int {
	maxProfit := 0
	buyDay := 0

	for day := 1; day < len(prices); day++ {
		if maxProfit < prices[day]-prices[buyDay] {
			maxProfit = prices[day] - prices[buyDay]
		}
		if prices[buyDay] > prices[day] {
			buyDay = day
		}
	}
	return maxProfit
}