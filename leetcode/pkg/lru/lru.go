package lru

import "fmt"

type LRUCache struct {
	mapping  map[string]*Node
	cache    *DoubleList
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	mapping := make(map[string]*Node)
	cache := NewDoubleList()

	return &LRUCache{
		mapping:  mapping,
		cache:    cache,
		capacity: capacity,
	}
}

// 提升为最近使用的节点
func (l *LRUCache) makeRecentlyUsed(key string) {
	n := l.mapping[key]
	l.cache.Remove(n)
	l.cache.AddLast(n)
}

// 添加为最近使用的节点
func (l *LRUCache) addRecently(key string, value int) {
	n := NewNode(key, value)
	l.cache.AddLast(n)
	l.mapping[key] = n
}

// 根据key删除节点
func (l *LRUCache) deleteKey(key string) {
	n := l.mapping[key]
	l.cache.Remove(n)
	delete(l.mapping, key)
}

// 删除最久未使用的元素
func (l *LRUCache) removeLastRecently() {
	n := l.cache.RemoveFirst()
	del_key := n.Key
	delete(l.mapping, del_key)
}

// 判断key是否存在
// 如果不存在，则返回-1
// 如果存在，则将当前key对应的节点提升为最近使用
// 并返回该节点的值
func (l *LRUCache) Get(key string) int {
	if _, ok := l.mapping[key]; !ok {
		return -1
	}
	l.makeRecentlyUsed(key)
	n := l.mapping[key]
	return n.Value
}

// 判断key是否存在
// 如果存在，则移除旧的节点
// 并把新的节点增加为最近使用的节点
// 如果不存在，则判断cache的长度是否等于capacity
// 如果等于，则删除最久没有使用的节点
// 如果不等于，则把新的节点增加为最近使用的节点
func (l *LRUCache) Put(key string, value int) {
	if _, ok := l.mapping[key]; ok {
		l.deleteKey(key)
		l.addRecently(key, value)
		return
	}

	if l.capacity == l.cache.Size() {
		l.removeLastRecently()
	}

	l.addRecently(key, value)
}

func (l *LRUCache) Show() {
	fmt.Println(l.cache)
}
