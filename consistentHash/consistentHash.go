package consistentHash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"sync"
)

const (
	defaultVirtualSpots = 128
)

type ConsistentHashRing struct {
	virtualSpots  int
	nodes         map[string]bool
	replicas      map[uint32]string //key: 虚拟节点id，value: node id
	hashSortNodes []uint32
	mu            sync.RWMutex
}

func NewConsistentHashRing(spots int) *ConsistentHashRing {
	if spots <= 0 {
		spots = defaultVirtualSpots
	}

	h := &ConsistentHashRing{
		virtualSpots:  spots,
		nodes:         make(map[string]bool, 0),
		replicas:      make(map[uint32]string, 0),
		hashSortNodes: make([]uint32, 0),
	}

	return h
}

func (h *ConsistentHashRing) AddNode(nodeKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.nodes[nodeKey]; ok {
		return
	}

	h.nodes[nodeKey] = true
	for i := 1; i <= h.virtualSpots; i++ {
		replicasKey := getReplicasKey(nodeKey, i)
		h.replicas[replicasKey] = nodeKey
		h.hashSortNodes = append(h.hashSortNodes, replicasKey)
	}

	sort.Slice(h.hashSortNodes, func(i, j int) bool {
		return h.hashSortNodes[i] < h.hashSortNodes[j]
	})
}

func (h *ConsistentHashRing) RemoveNode(nodeKey string) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.nodes[nodeKey]; !ok {
		return fmt.Errorf("%s not exists", nodeKey)
	}

	// 删除nodekey
	delete(h.nodes, nodeKey)

	// 删除虚拟节点key
	for i := 1; i < h.virtualSpots; i++ {
		replicasKey := getReplicasKey(nodeKey, i)
		delete(h.replicas, replicasKey)
	}

	// 移除旧的hashSortNodes， 并对新对hashSortNodes排序
	h.hashSortNodes = nil
	for k := range h.replicas {
		h.hashSortNodes = append(h.hashSortNodes, k)
	}

	sort.Slice(h.hashSortNodes, func(i, j int) bool {
		return h.hashSortNodes[i] < h.hashSortNodes[j]
	})

	return nil
}

func (h *ConsistentHashRing) Get(key string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.nodes) == 0 {
		return "", fmt.Errorf("Empty Nodes")
	}

	index := h.searchNearByIndex(key)
	nodeKey := h.replicas[h.hashSortNodes[index]]

	return nodeKey, nil
}

func (h *ConsistentHashRing) searchNearByIndex(key string) int {
	hk := hashKey(key)

	index := sort.Search(len(h.hashSortNodes), func(i int) bool {
		return h.hashSortNodes[i] >= hk
	})

	if index >= len(h.hashSortNodes) {
		index = 0
	}
	return index
}

func (h *ConsistentHashRing) Show() {
	fmt.Printf("nodes: %v\nreplicas: %v\nhashSortNodes: %v\n", h.nodes, h.replicas, h.hashSortNodes)
}

func getReplicasKey(nodeKey string, i int) uint32 {
	return hashKey(fmt.Sprintf("%s#%d", nodeKey, i))
}

func hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
