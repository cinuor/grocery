package stack

import (
	"fmt"
	"testing"
)

func TestValidBrackets(t *testing.T) {

	r := ValidBrackets("[[[]]]]")

	fmt.Println(r)
}

func TestBaseBallRaces(t *testing.T) {
	result := BaseBallRace([]string{"5", "2", "C", "D", "+"})

	fmt.Println(result)
}

func TestMyqueue(t *testing.T) {
	q := NewMyQueue()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
}

func TestSimpilfyPath(t *testing.T) {
	resp := SimplifyPath("/a/./b/../../c/")
	fmt.Println(resp)
}

func TestEvalRPN(t *testing.T) {
	r := EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	fmt.Println(r)
}

func TestTrap(t *testing.T) {
	// arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	arr := []int{2, 1, 0, 3}
	ans := Trap(arr)

	fmt.Println(ans)
}
