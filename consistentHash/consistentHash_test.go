package consistentHash

import (
	"fmt"
	"testing"
)

func Test_hashKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test cast 1",
			args: args{
				key: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashKey(tt.args.key)
			fmt.Println(got)
		})
	}
}

func TestConsistentHashRing_AddNode(t *testing.T) {

	h := NewConsistentHashRing(8)
	tests := []struct {
		name    string
		nodeKey string
	}{
		// TODO: Add test cases.
		{
			name:    "test case 1",
			nodeKey: "1000",
		},
		{
			name:    "test case 2",
			nodeKey: "2000",
		},
		{
			name:    "test case 1",
			nodeKey: "3000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.AddNode(tt.nodeKey)
		})
	}

	h.Show()
	h.RemoveNode("3000")
	h.Show()

	key, _ := h.Get("a")
	fmt.Println(key)
}
