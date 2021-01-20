package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	n := 0
	for x > n {
		n = n*10 + x%10
		x /= 10
	}
	return x == n || x == n/10
}

func main() {
	x := 1221

	fmt.Println(isPalindrome(x))
}
