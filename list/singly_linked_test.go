package main

import (
	"github.com/isdamir/gotype"
	"testing"
)

func CreateNodeT() (l1, l2 *gotype.LNode) {
	l1 = &gotype.LNode{}
	l2 = &gotype.LNode{}

	cur := l1
	for i := 1; i < 7; i++ {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}

	cur = l2
	for i := 9; i > 4; i-- {
		cur.Next = &gotype.LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
	return
}

func TestAdd(t *testing.T) {
	l1, l2 := CreateNodeT()
	gotype.PrintNode("l1: ", l1)
	gotype.PrintNode("l2: ", l2)
	addResult := Add(l1, l2)
	gotype.PrintNode("result: ", addResult)
}

func TestFindLastK(t *testing.T) {

	head1 := &gotype.LNode{}
	gotype.CreateNode(head1, 8)
	head2 := &gotype.LNode{}
	gotype.CreateNode(head2, 8)
	head3 := &gotype.LNode{}
	gotype.CreateNode(head3, 2)
	type args struct {
		head *gotype.LNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"FindLastK_OK", args{head: head1, k: 3}, 5},
		{"FindLastK_Short", args{head: head2, k: 8}, 3},
		{"FindLastK_nil", args{head: head3, k: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := FindLastK(tt.args.head, tt.args.k); err == nil && got.Data != tt.want {
				t.Errorf("FindLastK() = %v, want %v", got.Data, tt.want)
			}
		})
	}
}
