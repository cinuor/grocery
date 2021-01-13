package main

import (
	"fmt"
)

// 自顶向下，备忘录做法
// 定义：dp(n) 表示最少需要dp(n)个硬币才能凑成面值n
func CoinChange(coins []int, amount int) int {
	m := make(map[int]int, 0)
	return dp(coins, amount, m)
}

func dp(coins []int, n int, m map[int]int) int {
	if val, ok := m[n]; ok {
		return val
	}

	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	res := 2147483647
	for _, c := range coins {
		sub := dp(coins, n-c, m)
		if sub == -1 {
			continue
		}
		res = min(res, 1+sub)
	}

	if res == 2147483647 {
		m[n] = -1
	} else {
		m[n] = res
	}

	return m[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 自底向上, dp table做法
// 定义：dp[n]个硬币才成凑成面值n
// 也就是说, dp[i] = dp[i-c] + 1 |i-c > 0
func CoinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		if i == 0 {
			dp[i] = 0
			continue
		}
		dp[i] = amount + 1
	}

	fmt.Println(dp)

	for i := 0; i < len(dp); i++ {
		for _, c := range coins {
			if (i - c) < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}

	fmt.Println(dp)
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{3, 5, 7}
	amount := 4
	r := CoinChangeDP(coins, amount)
	fmt.Println(r)
}
