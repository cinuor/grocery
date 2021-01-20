package binarysearch

import (
	"testing"
)

func TestBasicBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "testcase1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7},
				target: 3,
			},
			want: 2,
		},
		{
			name: "testcase2",
			args: args{
				nums:   []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 7},
				target: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BasicBinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BasicBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftBoundaryBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "testcase1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7},
				target: 3,
			},
			want: 2,
		},
		{
			name: "testcase2",
			args: args{
				nums:   []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 7},
				target: 7,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LeftBoundaryBinarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("LeftBoundaryBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
