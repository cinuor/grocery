package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			nums: []int{1, 4, 2, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.nums)
			fmt.Println(tt.nums)
		})
	}
}
