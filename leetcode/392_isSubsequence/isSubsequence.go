package main

func isSubsequence(s string, t string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	i := 0
	for _, v := range t {
		if s[i] == byte(v) {
			i++
			if i == len(s) {
				return true
			}
		}
	}

	return false

}
