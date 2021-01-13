package main

import "fmt"

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	res := dp[0]

	for i := 0; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//  2, 5, 3, 7, 101
	fmt.Println(lengthOfLIS(nums))
}
