package sort

import (
	"fmt"
	"testing"
)

func Test_countingSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
			args: args{
				nums: []int{9, 4, 6, 2, 8, 1, 6, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countingSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
