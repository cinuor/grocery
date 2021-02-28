package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				nums:  []int{9, 3, 5, 2, 6, 8, 1},
				start: 0,
				end:   6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.nums)
			QuickSort(tt.args.nums, tt.args.start, tt.args.end)
			fmt.Println(tt.args.nums)
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				nums:  []int{9, 3, 5, 2, 6, 8, 1},
				start: 0,
				end:   6,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.nums, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}

			fmt.Println(tt.args.nums)
		})
	}
}

func BenchmarkQuickSortWithAsc(b *testing.B) {
	times := 1
	b.ResetTimer()

	sortedNums := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		sortedNums[i] = i
	}

	for i := 0; i < times; i++ {
		QuickSort(sortedNums, 0, 99999)
	}
}

func BenchmarkQuickSortDesc(b *testing.B) {
	times := 1
	b.ResetTimer()

	sortedNums := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		sortedNums[i] = 100000 - i
	}

	for i := 0; i < times; i++ {
		QuickSort(sortedNums, 0, 99999)
	}
}

func TestFindK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				nums: []int{4, 2, 5, 12, 3},
				k:    1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindK() = %v, want %v", got, tt.want)
			}
		})
	}
}
