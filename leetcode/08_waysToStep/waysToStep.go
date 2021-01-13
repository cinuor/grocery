package main

func waysToStep(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	s1, s2, s3 := 1, 2, 4

	for i := 3; i < n; i++ {
		s1, s2, s3 = mod(s2), mod(s3), mod(s1+s2+s3)
	}

	return s3

}

func mod(n int) int {
	return n % 1000000007
}
