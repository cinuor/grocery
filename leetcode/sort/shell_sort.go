package sort

import "fmt"


func ShellSort(nums []int) {
	N := len(nums)
	if N < 2 {
		return
	}
	for gap := N/2; gap > 0; gap /= 2{
		for i := gap; i < N; i++ {
			fmt.Println(nums)
			for j := i; j-gap >= 0 && nums[j] < nums[j-gap]; j -= gap {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
			}
		}
	}
}
