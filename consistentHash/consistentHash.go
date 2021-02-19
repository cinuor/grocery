package consistentHash

import (
	"sort"
	"sync"
)

type Node struct {
	nodeKey   string
	spotValue int
}

type NodeArray []Node

func (p NodeArray) Len() int           { return len(p) }
func (p NodeArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }
func (p NodeArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p NodeArray) Sort()              { sort.Sort(p) }

type Weight uint

const (
	defaultVirtualSpots = 512
)

type ConsistentHashRing struct {
	virtualSpots int
	hashNodes    NodeArray
	nodes        map[string]Weight
	mu           sync.RWMutex
}

func NewConsistentHashRing(spots int) *ConsistentHashRing {
	if spots <= 0 {
		spots = defaultVirtualSpots
	}

	h := &ConsistentHashRing{
		virtualSpots: spots,
		hashNodes:    make(NodeArray, 0),
		nodes:        make(map[string]Weight, 0),
	}

	return h
}

func (h *ConsistentHashRing) Add() error {

}
