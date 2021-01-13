package main

import "fmt"

func maxTurbulenceSize(arr []int) int {
	/*
	 *     length := len(arr)
	 *     ans := 1
	 *     pre := 0
	 *
	 *     for i := 1; i < length; i++ {
	 *         c := compare(arr[i-1], arr[i])
	 *         if i == length-1 || c*compare(arr[i], arr[i+1]) != -1 {
	 *             if c != 0 {
	 *                 ans = max(ans)
	 *             }
	 *         }
	 *     }
	 *
	 *     return ans
	 */
	coms := make([]int, len(arr)-1)

	if len(coms) == 0 {
		return 1
	}

	for x, i := arr[0], 1; i < len(arr); i++ {
		coms[i-1] = compare(x, arr[i])
		x = arr[i]
	}

	fmt.Println(coms)

	var pre int = 0
	if coms[0] != 0 {
		pre = 1
	}
	res := 0

	for i := 1; i < len(coms); i++ {
		if coms[i] == 0 {
			pre = 0
		}

		if coms[i-1]*coms[i] == -1 {
			pre = max(pre, pre+1)
		}

		// pre = max(pre, pre+1)
		// } else if coms[i-1]*coms[i] == 1 {
		// pre = 1
		// } else {
		// pre = 0
		// }

		res = max(pre, res)
	}

	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func compare(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func main() {

	// nums := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	// nums := []int{4, 8, 12, 16}
	// nums := []int{100, 100, 100}
	nums := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}
	x := maxTurbulenceSize(nums)
	fmt.Println(x)
}
