package sort

func InsertSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > temp {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = temp
	}
}
