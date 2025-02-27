package main

import (
	"errors"
	"fmt"
	"log"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

type SinglyLinkedList[T any] struct {
	Head *Node[T]
}

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Head: nil,
	}
}

func (ll *SinglyLinkedList[T]) GetSize() int {
	size := 0
	curr := ll.Head
	for curr != nil {
		size++
		curr = curr.Next
	}

	return size
}

func (list *SinglyLinkedList[T]) Add(value T) {
	newNode := &Node[T]{
		Val:  value,
		Next: nil,
	}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *SinglyLinkedList[T]) RemoveFirstItem() error {
	if list.Head == nil || list.Head.Next == nil {
		return errors.New("cannot remove from an empty list")
	}

	current := list.Head.Next
	list.Head = nil
	list.Head = current

	return nil
}

func (list *SinglyLinkedList[T]) Remove(idx int) error {
	if list.Head == nil {
		return errors.New("list is empty")
	}

	if idx < 0 {
		return errors.New("index is out of bounds")
	}

	if idx == 0 {
		list.Head = list.Head.Next
		return nil
	}

	current := list.Head
	for i := 0; i < idx-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return errors.New("index out of bounds")
	}

	current.Next = current.Next.Next
	return nil
}

func (list *SinglyLinkedList[T]) RemoveLastItem() {
	current := list.Head

	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (list *SinglyLinkedList[T]) Print() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func addNode(head *ListNode, newNode *ListNode) {
	for {
		if head.Next == nil {
			head.Next = newNode
			break
		}

		head = head.Next
	}
}

func removeLastNode(head *ListNode) {
	if head == nil {
		return
	}

	for {
		if head.Next.Next == nil {
			head.Next = nil
			break
		}

		head = head.Next
	}
}

func removeNodeAtIndex(head *ListNode, index int) {
	if head == nil {
		return
	}

	if index == 0 {
		return
	}

	prev := head
	current := head.Next
	i := 1

	for current != nil {
		if i == index {
			prev.Next = current.Next
			break
		}

		prev = current
		current = current.Next
		i++
	}
}

func printLinkedList(head *ListNode) {
	for {
		if head == nil {
			break
		}

		fmt.Println(head.Val)
		head = head.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	cur := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		// add the sum to val
		// [2, 4, 3] [5, 6, 4]
		s := carry
		if l1 != nil {
			s += l1.Val
		}
		if l2 != nil {
			s += l2.Val
		}
		// [7]

		carry = s / 10
		// 7 / 10 = 0.7
		cur.Next = &ListNode{Val: s % 10, Next: nil}
		// Val: 7 % 10 = 7

		cur = cur.Next
		// Traverse to next node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummy.Next
}

func main() {
	singlyll := NewSinglyLinkedList[int]()

	singlyll.Add(1)
	singlyll.Add(2)
	singlyll.Add(3)
	singlyll.Add(4)
	singlyll.Add(5)
	singlyll.Add(6)

	err := singlyll.RemoveFirstItem()
	if err != nil {
		log.Fatal(err)
	}
	singlyll.RemoveFirstItem()
	singlyll.RemoveLastItem()
	llSize := singlyll.GetSize()
	fmt.Println("size: ", llSize)

	singlyll.Print()
}
