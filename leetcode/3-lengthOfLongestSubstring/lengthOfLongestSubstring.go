package main

import (
    "fmt"
    "os"
)

func lengthOfLongestSubstring(s string) int {
    maxLength, i := 0,0
    sc := make(map[rune]int)

    for j, ch := range([]rune(s)){
        if k, ok := sc[ch]; ok {
            // if j - k >= maxLength {
                // maxLength = j - k
            // }
        }
        sc[ch] = j
        if j - i + 1 > maxLength {
            maxLength = j - i + 1
        }
    }

    return maxLength
}

func main() {
    s := string(os.Args[1])

    results := lengthOfLongestSubstring(s)

    fmt.Printf("%d\n", results)

}
