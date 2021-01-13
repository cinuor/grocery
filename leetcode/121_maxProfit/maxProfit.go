package main

import "fmt"

// dp[i] = max(dp[i-1], prices[i]-minPrice)

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}

	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	x := maxProfit(prices)
	fmt.Println(x)
}
