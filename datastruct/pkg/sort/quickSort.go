package sort

func QuickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(nums, start, end)
	QuickSort(nums, start, pivot-1)
	QuickSort(nums, pivot+1, end)

}

func partition(nums []int, start, end int) int {
	pivot := nums[end]

	l := start
	r := end

	for l != r {
		for l < r && nums[l] < pivot {
			l++
		}

		for l < r && nums[r] >= pivot {
			r--
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[l], nums[end] = nums[end], nums[l]
	return l
}

func partitionDesc(nums []int, start, end int) int {
	pivot := nums[end]

	l := start
	r := end

	for l != r {
		for l < r && nums[l] > pivot {
			l++
		}

		for l < r && nums[r] <= pivot {
			r--
		}

		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	nums[l], nums[end] = nums[end], nums[l]
	return l
}

//找出第K大的数, 0 < k <= len(nums)
func FindK(nums []int, k int) int {
	start, end := 0, len(nums)-1
	p := partitionDesc(nums, start, end)
	for p+1 != k {
		if p+1 < k {
			p = partitionDesc(nums, p+1, end)
		} else {
			p = partitionDesc(nums, start, p-1)
		}
	}
	return nums[p]
}
