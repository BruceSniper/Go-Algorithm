package Leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	Min := prices[0]
	Max := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-Min > Max {
			Max = prices[i] - Min
		}
		if prices[i] < Min {
			Min = prices[i]
		}
	}
	return Max
}
