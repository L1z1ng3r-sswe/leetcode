func maxProfit(prices []int) int {
	cash, hold := 0, -prices[0]
	prevCash := 0

	for i := 1; i < len(prices); i++ {
		temp := cash
		cash = max(cash, hold+prices[i])
		hold = max(hold, prevCash-prices[i])
		prevCash = temp
	}

	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// time:  O(N)
// space: O(1)