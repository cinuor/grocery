package main

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {

	s := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := s[target-v]; ok {
			return []int{j, i}
		}

		s[v] = i
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 2, 15}
	var target int = 4

	fmt.Println(TwoSum(nums, target))
}
