package _2MaxProfit

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}

		if prices[i]-buyPrice > profit {
			profit = prices[i] - buyPrice
		}
	}
	return profit
}
