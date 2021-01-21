package lru

type Node struct {
	Key   string
	Value int
	Next  *Node
	Prev  *Node
}

func NewNode(key string, value int) *Node {
	return &Node{Key: key, Value: value}
}
