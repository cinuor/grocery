package lru

import (
	"testing"
)

func TestLRUCache_Put(t *testing.T) {

	l := NewLRUCache(4)

	type args struct {
		nodes []*Node
	}
	tests := []struct {
		name string
		l    *LRUCache
		args args
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
			l:    l,
			args: args{
				nodes: []*Node{
					NewNode("a", 1),
					NewNode("b", 1),
					NewNode("c", 1),
					NewNode("d", 1),
					NewNode("e", 1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.args.nodes {
				tt.l.Put(tt.args.nodes[i].Key, tt.args.nodes[i].Value)
			}
		})
	}

	l.Show()

}

func TestLRUCache_Get(t *testing.T) {
	l := NewLRUCache(4)

	type args struct {
		key string
	}
	tests := []struct {
		name string
		l    *LRUCache
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
			l:    l,
			args: args{
				key: "a",
			},
			want: -1,
		},
		{
			name: "testcase 2",
			l:    l,
			args: args{
				key: "b",
			},
			want: 1,
		},
		{
			name: "testcase 3",
			l:    l,
			args: args{
				key: "c",
			},
			want: 1,
		},
	}

	nodes := []*Node{
		NewNode("a", 1),
		NewNode("b", 1),
		NewNode("c", 1),
		NewNode("d", 1),
		NewNode("e", 1),
	}

	for i := range nodes {
		l.Put(nodes[i].Key, nodes[i].Value)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})

		l.Show()
	}
}
