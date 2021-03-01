package sort

import (
	"fmt"
	"testing"
)

func Test_bucketSort(t *testing.T) {
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
				nums: []int{31, 45, 67, 23, 91, 100, 66, 55, 22, 90, 53, 67, 88, 345, 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.nums)
			bucketSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
