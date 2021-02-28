package queue

import "fmt"

type CircularQueue struct {
	items []int
	n     int
	head  int
	tail  int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, capacity, capacity),
		n:     capacity,
		head:  0,
		tail:  0,
	}
}

func (q *CircularQueue) Enqueue(item int) bool {
	if (q.tail+1)%q.n == q.head {
		return false
	}

	q.items[q.tail] = item
	q.tail += 1
	return true
}

func (q *CircularQueue) Dequeue() (int, bool) {
	if q.head == q.tail {
		return 0, false
	}

	item := q.items[q.head]
	q.items[q.head] = 0
	q.head = (q.head + 1) % q.n
	return item, true
}

func (q *CircularQueue) Show() {
	for i := 0; i < q.n; i++ {
		fmt.Printf("%v ", q.items[i])
	}
	fmt.Println("")
}
