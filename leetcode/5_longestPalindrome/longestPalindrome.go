package main

import "fmt"

func longestPalindrome(s string) string {

	var s1, s2, res string

	for i := 0; i < len(s); i++ {
		s1 = palindrome(s, i, i)
		s2 = palindrome(s, i, i+1)

		res = Max(res, Max(s1, s2))
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for {
		if l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		} else {
			break
		}
	}
	return s[l+1 : r]
}

func Max(a string, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func main() {
	s := "babad"

	x := longestPalindrome(s)
	fmt.Println(x)
	/*
	 * x := palindrome(s, 0, 0)
	 * fmt.Println(x)
	 * y := palindrome(s, 0, 1)
	 * fmt.Println(y)
	 */

}
