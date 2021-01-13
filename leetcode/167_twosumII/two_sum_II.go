package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers) - 1
	for ;low < high; {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low+1, high+1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}

