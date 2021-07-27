package main

import (
	"errors"
	"fmt"
	"sync"
)

type Node struct {
	value interface{}
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

type SinglyLinkedList struct {
	head *Node
	len  int
	lock sync.RWMutex
}

func NewSinglyLinked() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (s *SinglyLinkedList) Append(n *Node) {
	if s.head == nil {
		s.head = n
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}

	s.len++
}

func (s *SinglyLinkedList) Prepend(n *Node) {
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.len++
}

func (s *SinglyLinkedList) GetHead() *Node {
	return s.head
}

func (s *SinglyLinkedList) GetBack() *Node {
	current := s.head
	if current != nil && current.next != nil {
		current = current.next
	}
	return current
}

func (s *SinglyLinkedList) Traverse() error {
	if s.head == nil {
		return errors.New("TraverseError: List is empty")
	}

	current := s.head
	for current != nil {
		fmt.Printf(" %v ->", current.value)
		current = current.next
	}

	return nil
}

func (s *SinglyLinkedList) Reverse() {
	curr := s.head
	var prev *Node
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	s.head = prev
}

func main() {
	l := NewSinglyLinked()
	l.Append(&Node{
		value: "a",
	})
	l.Append(&Node{
		value: "b",
	})
	l.Append(&Node{
		value: 1111,
	})
	l.Prepend(&Node{
		value: 1})

	if err := l.Traverse(); err != nil {
		fmt.Println(err)
	}

	l.Reverse()
	fmt.Println()
	if err := l.Traverse(); err != nil {
		fmt.Println(err)
	}
}
