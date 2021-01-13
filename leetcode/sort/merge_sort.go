package sort

import (
	// "fmt"
	// "math"
)

func merge(nums []int, temp []int, low int, mid int, high int) {
	for i, val := range nums {
		temp[i] = val
	}

	i := low
	j := mid + 1

	for k := low; k <= high ; k++ {
		if i > mid {
			nums[k] = temp[j]
			j++
		} else if j > high {
			nums[k] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			nums[k] = temp[i]
			i++
		} else {
			nums[k] = temp[j]
			j++
		}
	}
}

func mergeSort(nums []int, temp []int, low int, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	mergeSort(nums, temp, low, mid)
	mergeSort(nums, temp, mid+1, high)
	merge(nums, temp, low, mid, high)
}

func MergeSortUpToDown(nums []int) {
	temp := make([]int, len(nums), len(nums))
	mergeSort(nums, temp, 0, len(nums)-1)
}

// func MergeSortDownToUp(nums []int) {
	// length := len(nums)
	// // temp := make([]int, length, length)

	// maxRound := int(math.Sqrt(float64(length)))+1
	// fmt.Println(maxRound)
	// for round := 0; round < maxRound; round += 1{
		// sep := int(math.Pow(2, float64(round)))
		// for low := 0; low < low + sep; low = sep + sep  {
			// fmt.Println(low, low+sep-1, low+sep)
			// // merge(nums, temp, low, low+sep-1, low+sep)
		// }
	// }
// }
