package slidewindow

/*
 * func slidingWindow(s string, t string) {
 *
 *     rs := []rune(s)
 *     rt := []rune(t)
 *
 *     window := make(map[rune]int)
 *
 *     left, right := 0, 0 // 两个指针，[left, right)
 *     for right < len(rs) {
 *         c := rs[right]
 *         right++
 *
 *         // 添加字符到window
 *         // 如何c已经存在则计数加一，如果不存在则添加
 *         if count, ok := window[c]; ok {
 *             window[c] = count + 1
 *         } else {
 *             window[c] = 1
 *         }
 *
 *         // TODO: 具体的业务操作
 *
 *         //判断左侧窗口是否需要收缩
 *         for condition {
 *             c := s[left]
 *             left++
 *
 *             // TODO: 具体的业务操作
 *         }
 *
 *     }
 * }
 */

/*
 * func minWindow(s string, t string) string {
 * }
 */

// "abcabcbb"
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := window[c]; ok {
			window[c] += 1
		} else {
			window[c] = 1
		}

		for window[c] > 1 {
			d := s[left]
			window[d] -= 1
			left++
		}

		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
