package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_Show(t *testing.T) {
	q := NewCircularQueue(8)
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				q.Enqueue(i + 1)
			}

			x, _ := q.Dequeue()
			fmt.Println(x)

			q.Show()
		})
	}
}
