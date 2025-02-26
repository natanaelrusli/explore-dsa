package main

import (
	"fmt"
)

func main() {
	// Create a new stack
	stack := NewStack()

	// Let's demonstrate stack operations
	fmt.Println("Pushing numbers 1, 2, and 3 onto the stack")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Printf("Stack size: %d \n", stack.Size())

	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Top element (peek): %d \n", top)
	}

	fmt.Println("\\nPopping all elements:")
	for !stack.IsEmpty() {
		if item, err := stack.Pop(); err == nil {
			fmt.Printf("Popped: %d \n", item)
		}
	}

	linkedListStack := StackNode{}
	linkedListStack.Push(1)
	linkedListStack.Push(2)
	linkedListStack.Push(3)
	linkedListStack.Push(4)

	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()
	linkedListStack.Pop()

	fmt.Printf("Stack size: %d \n", linkedListStack._size)
}
