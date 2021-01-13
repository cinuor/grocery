package array

import (
// "fmt"
// "math"
)

// 42.接雨水
// 思路：双指针

func Trap(arr []int) int {
	l, r := 0, len(arr)-1
	lmax, rmax, ans := 0, 0, 0

	for l < r {
		if arr[l] < arr[r] { // 左边的墙比右边矮，所以以左边的墙高为准
			if arr[l] > lmax {
				lmax = arr[l]
			}
			ans += lmax - arr[l]
			l++
		} else {
			if arr[r] > rmax {
				rmax = arr[r]
			}
			ans += rmax - arr[r]
			r--
		}
	}

	return ans
}

/*
 * 错误的想法
 * func Trap(arr []int) int {
 *     var a, b, ans int = 0, 0, 0
 *
 *     for b < len(arr) {
 *         if arr[a] <= arr[b] && a < b {
 *             h := int(math.Min(float64(arr[a]), float64(arr[b])))
 *             w := b - a - 1
 *             total := w * h
 *             for i := a + 1; i <= b-1; i++ {
 *                 total -= arr[i]
 *             }
 *             ans += total
 *             fmt.Println(a, b)
 *             a = b
 *             continue
 *         }
 *
 *         b += 1
 *
 *     }
 *
 *     return ans
 * }
 */
