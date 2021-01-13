package main

import (
	"fmt"
	"math/rand"
)

const (
	MaxLevel = 32
	P        = 0.25
)

type node struct {
	key   int
	value interface{}
	next  []*node
}

func randomLevel() int {
	level := 1
	for rand.Float32() < P && level < MaxLevel {
		level++
	}

	return level
}

func NewNode(key int, value interface{}, level int) *node {
	return &node{key, value, make([]*node, level)}
}

type SkipList struct {
	head  *node
	level int
}

func NewSkipList() *SkipList {
	return &SkipList{head: NewNode(0, nil, MaxLevel)}
}

func (s *SkipList) Search(key int) (*node, bool) {
	cur := s.head

	for i := s.level - 1; i >= 0; i-- { // 每一层的头节点开始
		for cur.next[i] != nil && cur.next[i].key < key { // key小于目标值的节点，就继续往下迭代
			cur = cur.next[i]
		}

		// 此时，节点的下一个节点为空或者key >= 目标值
		// 换层
	}

	// 直到到达最后一层, 找到目标节点
	// 因为第0层的时候，所有skiplist退化成了单链表
	// 所以目标节点能通过cur.next[0]索引得到
	cur = cur.next[0]

	if cur != nil && cur.key == key {
		return cur, true
	}

	return nil, false
}

func (s *SkipList) Insert(key int, value interface{}) *node {
	cur := s.head
	update := make([]*node, MaxLevel)

	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key < key {
			cur = cur.next[i]
		}

		// 记录当前层第一个不满足要求的节点（它的next可能为空，或者它的key >= 目标key）
		// 插入位置就在该节点之后
		update[i] = cur
	}

	// 找到目标节点
	cur = cur.next[0]

	if cur != nil && cur.key == key {
		cur.value = value
		return cur
	}

	level := randomLevel()
	if level > s.level {
		level = s.level + 1
		update[s.level] = s.head
		s.level = level
	}

	n := NewNode(key, value, level)

	for i := 0; i < level; i++ {
		n.next[i] = update[i].next[i]
		update[i].next[i] = n
	}

	return n
}

func (s *SkipList) Delete(key int) *node {
	update := make([]*node, MaxLevel)
	cur := s.head

	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].key < key {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	cur = cur.next[0]

	if cur != nil && cur.key == key {
		for i := 0; i < s.level; i++ {
			if update[i].next[i] != cur {
				return nil
			}
			update[i].next[i] = cur.next[i]
		}
	}
	return cur
}

func (s *SkipList) Print() {
	for i := s.level - 1; i >= 0; i-- {
		cur := s.head
		fmt.Printf("head")
		for cur.next[i] != nil {
			fmt.Printf("->%d", cur.next[i].key)
			cur = cur.next[i]
		}
		fmt.Printf("\n")
	}
}

func main() {
	l := NewSkipList()

	n := 10

	for i := 0; i < n; i++ {
		l.Insert(i, i)
	}

	l.Print()

	x, ok := l.Search(90)
	if ok {
		fmt.Println(x.key)
	}
}
