package sort

func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	flag := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}

		if !flag {
			break
		}

	}
}
