package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i]) // 状态转移方程
		ans = max(ans, nums[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	x := maxSubArray(nums)
	fmt.Println(x)
}
