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
