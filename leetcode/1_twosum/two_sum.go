package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    store := make(map[int]int, len(nums))
    for i, v := range nums {

        if j, ok := store[target - v]; ok {
            return []int{j, i}
        }
        store[v] = i
    }

    return nil
}

func main() {
    var nums = []int{2,7,2,15}
    var target int = 4

    fmt.Println(twoSum(nums, target))
}
