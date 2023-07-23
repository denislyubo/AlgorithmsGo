package stockmaxprofit

import "fmt"

func maxProfit(stock []int) int {
	size := len(stock)
	buy := 0
	sell := 0
	curMin := 0
	currProfit := 0
	maxProfit := 0
	for i := 0; i < size; i++ {
		if stock[i] < stock[curMin] {
			curMin = i
		}
		currProfit = stock[i] - stock[curMin]
		if currProfit > maxProfit {
			buy = curMin
			sell = i
			maxProfit = currProfit
		}
	}

	fmt.Println("Purchase day is- ", buy, " at price ", stock[buy])
	fmt.Println("Sell day is- ", sell, " at price ", stock[sell])
	fmt.Println("Max Profit :: ", maxProfit)

	return maxProfit
}
