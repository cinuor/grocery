package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{8, 3, 5, 2, 6, 1, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := MergeSort(tt.args.nums)
			fmt.Println(resp)
		})
	}
}
