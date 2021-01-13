package main

import (
	"fmt"
	u "leetcode/pkg/utils"
)

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	/*
	 *     dp := make([]int, n)
	 *     dp[0] = nums[0]
	 *     dp[1] = nums[1]
	 *
	 *     for i := 2; i < n; i++ {
	 *         dp[i] = u.Max(dp[i-2]+nums[i], dp[i-1])
	 *     }
	 *
	 *     return dp[n-1]
	 */

	before := nums[0]
	prev := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		before, prev = prev, u.Max(before+nums[i], prev) // 使用两个变量来记录前前一个和前一个的盗窃最大值。
	}

	return prev
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nums := []int{1, 2, 3, 1}
	nums := []int{2, 7, 9, 3, 1}
	r := rob(nums)
	fmt.Println(r)
}
