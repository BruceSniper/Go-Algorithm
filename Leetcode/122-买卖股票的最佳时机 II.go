package Leetcode

func maxProfit(prices []int) int {
	money := 0
	now := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > now {
			money += prices[i] - now
		}
		now = prices[i]
	}
	return money
}
