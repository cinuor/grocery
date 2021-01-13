package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	res := nums[0]
	former := nums[0]
	for i := 1; i < n; i++ {
		former = max(former+nums[i], nums[i])
		res = max(res, former)
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
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(nums)
	fmt.Println(r)
	fmt.Println(nums)
}
