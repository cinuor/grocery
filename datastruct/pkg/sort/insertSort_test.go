package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				nums: []int{9, 5, 3, 2, 8, 7, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
