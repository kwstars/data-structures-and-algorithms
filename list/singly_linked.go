package main

import (
	"errors"
	"github.com/isdamir/gotype"
)

// Reverse Linklist
func Reverse(node *gotype.LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var (
		pre *gotype.LNode
		cur *gotype.LNode
	)

	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre

}

// RemoveDup Remove duplicate items from an unordered linklist
func RemoveDup(head *gotype.LNode) {
	if head == nil || head.Next == nil {
		return
	}

	outerCur := head.Next
	var (
		innerCur *gotype.LNode
		innerPre *gotype.LNode
	)

	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

// Add Summing two lined list
func Add(l1 *gotype.LNode, l2 *gotype.LNode) (head *gotype.LNode) {
	if l1 == nil || l1.Next == nil {
		return l2
	}

	if l2 == nil || l2.Next == nil {
		return l1
	}

	c := 0
	sum := 0
	p1 := l1.Next
	p2 := l2.Next
	resultHead := &gotype.LNode{}
	p := resultHead

	for p1 != nil && p2 != nil {
		p.Next = &gotype.LNode{}
		sum = p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum % 10
		c = sum / 10
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil {
		for p2 != nil {
			p.Next = &gotype.LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}

	if p2 == nil {
		for p1 != nil {
			p.Next = &gotype.LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}

	if c == 1 {
		p.Next = &gotype.LNode{}
		p.Next.Data = 1
	}

	return resultHead
}

func FindLastK(head *gotype.LNode, k int) (*gotype.LNode, error) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	var (
		i    = 0
		slow = head.Next
		fast = head.Next
	)

	for ; i < k && fast != nil; i++ {
		fast = fast.Next
	}

	if i < k {
		return nil, errors.New("linkedList too short")
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if slow == nil {
		return nil, errors.New("Not found")
	}

	return slow, nil
}
