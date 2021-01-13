package main

import "fmt"

// 斐波那契数列数列的变种
// 找出dp[i]的关键点是: 当前节点的最小花费，可以从前一级台阶跨一格，或者从前两级台阶垮个大劈叉。
func minCostClimbingStairs(cost []int) int {
	prev, cur := 0, 0
	n := len(cost)

	for i := 2; i <= n; i++ {
		prev, cur = cur, min(cur+cost[i-1], prev+cost[i-2])
	}

	return cur
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	cost := []int{10, 15, 20}

	r := minCostClimbingStairs(cost)
	fmt.Println(r)
}
