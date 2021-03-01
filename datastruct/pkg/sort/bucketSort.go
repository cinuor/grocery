package sort

import (
	"fmt"
	"math"
)

const capacity = 10

func bucketSort(nums []int) {

	// 1. 找出minValue和maxValue，确定小数组个数
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	bucketLen := (max-min)/capacity + 1
	bucket := make([][]int, bucketLen, bucketLen)

	for i := 0; i < len(nums); i++ {
		index := (nums[i] - min) / capacity
		fmt.Println(nums[i], index)
		bucket[index] = append(bucket[index], nums[i])
	}

	pos := 0
	for i := 0; i < bucketLen; i++ {
		if len(bucket[i]) == 0 {
			continue
		}

		InsertSort(bucket[i])
		for idx := range bucket[i] {
			nums[pos] = bucket[i][idx]
			pos++
		}
	}
}
