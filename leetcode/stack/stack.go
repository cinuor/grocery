package stack

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"sync"
)

type node struct {
	val  interface{}
	prev *node
}

type Stack struct {
	top    *node
	length int
	lock   sync.Mutex
}

func NewStack() *Stack {
	return &Stack{top: nil, length: 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.val
}

func (s *Stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &node{val: value, prev: s.top}
	s.top = n
	s.length++
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.top
	s.top = n.prev
	s.length--
	return n.val
}

func (s *Stack) MarshalJSON() ([]byte, error) {
	if s.length == 0 {
		return json.Marshal([]interface{}{})
	}

	var nodes []interface{}

	for i, head := 0, s.top; i < s.length && head != nil; i, head = i+1, head.prev {
		nodes = append(nodes, head.val)
	}

	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}

	return json.Marshal(nodes)

}

func (s *Stack) UnmarshalJSON(data []byte) error {
	var vals []interface{}

	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}

	for i := 0; i < len(vals); i++ {
		s.Push(vals[i])
	}

	return nil
}

// 20.有效的括号
func ValidBrackets(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 > 0 {
		return false
	}

	st := NewStack()
	enum := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			st.Push(v)
		} else {
			if st.Len() == 0 || st.Peek() != enum[v] {
				return false
			} else {
				st.Pop()
			}
		}
	}

	if st.Len() > 0 {
		return true
	}

	return true
}

// 682.棒球比赛
func BaseBallRace(s []string) int {
	st := NewStack()
	// var result int

	for _, r := range s {
		if r == "C" {
			st.Pop()
		} else if r == "D" {
			st.Push(2 * st.Peek().(int))
		} else if r == "+" {
			top := st.Peek()
			st.Push(top)
			st.Push(top.(int) + st.Peek().(int))
		} else {
			n, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			st.Push(n)
		}
	}

	var result int
	for i := 0; i < st.Len(); i++ {
		result += st.Peek().(int)
		st.Pop()
	}

	return result
}

// 71.简化路径
func SimplifyPath(s string) string {
	ret := []string{}

	for _, v := range strings.Split(s, "/") {
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(ret); l > 0 {
				ret = ret[:l-1]
			}
		default:
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}

// 150.逆波兰表达式求值
func EvalRPN(tokens []string) int {
	st := NewStack()

	for _, t := range tokens {
		switch t {
		case "+":
			num1 := st.Pop().(int)
			num2 := st.Pop().(int)
			st.Push(num2 + num1)
		case "-":
			num1 := st.Pop().(int)
			num2 := st.Pop().(int)
			st.Push(num2 - num1)
		case "*":
			num1 := st.Pop().(int)
			num2 := st.Pop().(int)
			st.Push(num2 * num1)
		case "/":
			num1 := st.Pop().(int)
			num2 := st.Pop().(int)
			st.Push(num2 / num1)
		default:
			num, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			st.Push(num)
		}
	}
	return st.Pop().(int)
}

// 42.接雨水
// 思路：单调递减栈
// 双指针见array/array.go:Trap
func Trap(arr []int) int {
	st := make([]int, 0)
	ans := 0

	for i := 0; i < len(arr); i++ {
		for len(st) != 0 && arr[i] > arr[st[len(st)-1]] {
			cur := st[len(st)-1]
			st = st[0 : len(st)-1]

			if len(st) == 0 {
				break
			}

			h := int(math.Min(float64(arr[st[len(st)-1]]), float64(arr[i]))) - arr[cur]
			w := i - st[len(st)-1] - 1
			ans += h * w
		}
		st = append(st, i)
	}
	return ans
}
