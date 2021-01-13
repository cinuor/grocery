package sort

func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i+1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		temp := nums[i]
		nums[i] = nums[min]
		nums[min] = temp
	}
}
