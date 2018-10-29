package main

import "fmt"

type node struct {
	value int
	next      *node
}

func main() {
	head := makeList()
	removedNode := remove(&head, 2)

	fmt.Printf("removed node :%+v\n", removedNode)
}

func remove(head **node, v int) *node {

	doublePointer := head

	for *doublePointer != nil && (**doublePointer).value != v  {
		doublePointer = &(*doublePointer).next
	}

	removedNode := *doublePointer
	*doublePointer = removedNode.next
	removedNode.next = nil

	return removedNode
}


func makeList() *node {
	var head *node

	for i := 5; i > 0; i-- {
		newNode := &node{i, head}
		head = newNode
	}

	return head
}

