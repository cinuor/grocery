package sort

import (
	// "fmt"
	"fmt"
	"testing"
)

func Test_Merge(t *testing.T) {
	// nums := []int{4, 5, 6, 7, 1, 2, 3}
	nums := []int{10, 3, 5, 2, 6, 7, 9}
	QuickSort(nums)
	fmt.Println(nums)
}
