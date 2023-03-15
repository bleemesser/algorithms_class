package datastructures

import (
	"fmt"
	"os"
)

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func makeNode(val int) *BinaryTreeNode {
	return &BinaryTreeNode{Val: val}
}

func (n *BinaryTreeNode) insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = makeNode(val)
		} else {
			n.Left.insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = makeNode(val)
		} else {
			n.Right.insert(val)
		}
	}
}

func BinaryTree(vals []int) *BinaryTreeNode {
	root := makeNode(vals[0])
	for _, val := range vals[1:] {
		root.insert(val)
	}
	return root
}

func (n *BinaryTreeNode) PrintBinaryTree(ns int, ch rune) {
	if n == nil {
		return
	}
	if ns == 0 {
		fmt.Fprint(os.Stdout, "\n")
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(os.Stdout, " ")
	}
	fmt.Fprintf(os.Stdout, "%c:%v\n", ch, n.Val)
	n.Left.PrintBinaryTree(ns+2, 'L')
	n.Right.PrintBinaryTree(ns+2, 'R')
}

func (n *BinaryTreeNode) ExportToSlice() []int {
	if n == nil {
		return nil
	}
	var vals []int
	vals = append(vals, n.Left.ExportToSlice()...)
	vals = append(vals, n.Val)
	vals = append(vals, n.Right.ExportToSlice()...)
	return vals
}
