package datastructures

import (
	"fmt"
)

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

func makeListItem(val int) *LinkedListNode {
	return &LinkedListNode{val, nil}
}

func (n *LinkedListNode) Insert(val int) {
	if n.Next == nil {
		n.Next = makeListItem(val)
	} else {
		n.Next.Insert(val)
	}
}

func (n *LinkedListNode) Print() {
	arr := []int{}
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}
	fmt.Println(arr)
}

func (n *LinkedListNode) Delete(val int) {
	if n.Next == nil {
		return
	}
	if n.Next.Val == val {
		n.Next = n.Next.Next
	} else {
		n.Next.Delete(val)
	}
}

func (n *LinkedListNode) Export() []int {
	arr := []int{}
	for n != nil {
		arr = append(arr, n.Val)
		n = n.Next
	}
	return arr
}

func LinkedList(arr []int) *LinkedListNode {
	if len(arr) == 0 {
		return nil
	}
	head := makeListItem(arr[0])
	for _, v := range arr[1:] {
		head.Insert(v)
	}
	return head
}
