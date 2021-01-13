package array

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ans := Trap(arr)
	fmt.Println(ans)
}
