package main

import "log"

// linked list
 type Node struct {
	data int
	next *Node
 }

 type StackNode struct {
	head *Node
	_size int // private size property
 }

 // function to push data onto the top node of the stack
 func (s *StackNode) Push(data int) {
	var newNode Node
	newNode.data = data
	if s.head == nil {
		s.head = &Node{data, nil}
	} else {
		// create a new node with next pointing to head
		newNode := &Node{data, s.head}
		s.head = newNode
	}

	s._size++
	if s._size == 1 {
		s.head = &newNode
	}
 }

 func (s *StackNode) adjustSizeForNewHead() {
	if s.head == nil && s._size > 0 {
		// when the head is added but size has been reduced by a Pop() operation
		s._size--
	}
 }

 // function to remove data from the top node of the stack
 // which is the last element added
 func (s *StackNode) Pop() int {
	if s.head == nil {
		log.Println("Out of bound")
		return -1
	}

	data := s.head.data
	s.head = s.head.next
	s._size--

	return data
 }

 // this function allows you to view the top element of the stack without removing it
 func (s *StackNode) Peek() int {
	if s.head == nil {
		log.Println("Out of bound")
		return -1
	}

	return s.head.data
 }

 func (s *StackNode) Size() int {
	return s._size
 }

 func (s *StackNode) GetTopElements() []int {
	var elements []int
	current := s.head

	for current != nil {
		elements = append(elements, current.data)
		current = current.next
	}

	return elements
 }

 func (s *StackNode) IsEmptyAtStart() bool {
	if s.head == nil && s._size > 0 {
		return true
	}
	return false
 }
