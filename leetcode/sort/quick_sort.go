package sort

func partition(nums []int, low int, high int) int {
	pivot := low
	i, j := low, high
	for i != j {
		for nums[i] < nums[pivot] && i < j {
			i++
		}

		for nums[j] > nums[pivot] && i < j {
			j--
		}
		if (i < j) {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[pivot], nums[i]= nums[i], nums[pivot]

	return pivot
}

func quickSort(nums []int, low int, high int) {
	if low < high {
		part := partition(nums, low, high)
		quickSort(nums, low, part - 1)
		quickSort(nums, part + 1, high)
	}
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}
