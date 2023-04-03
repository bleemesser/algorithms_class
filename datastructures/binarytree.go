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

func (n *BinaryTreeNode) Insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = makeNode(val)
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = makeNode(val)
		} else {
			n.Right.Insert(val)
		}
	}
}

func BinaryTree(vals []int) *BinaryTreeNode {
	root := makeNode(vals[0])
	for _, val := range vals[1:] {
		root.Insert(val)
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

func FindMin(n *BinaryTreeNode) int {
	if n == nil {
		return 0
	}
	if n.Left == nil {
		return n.Val
	}
	return FindMin(n.Left)
}

func FindVal(n *BinaryTreeNode, val int) bool {
	if n == nil {
		return false
	}
	if n.Val == val {
		return true
	}
	if val < n.Val {
		return FindVal(n.Left, val)
	}
	return FindVal(n.Right, val)
}

func (n *BinaryTreeNode) Delete(val int) *BinaryTreeNode {
	if n == nil {
        return nil
    }
    if val < n.Val {
        n.Left = n.Left.Delete(val)
        return n
    }
    if val > n.Val {
        n.Right = n.Right.Delete(val)
        return n
    }
    // If the node to be deleted has no children, just return nil
    if n.Left == nil && n.Right == nil {
        return nil
    }
    // If the node to be deleted has only one child, return that child
    if n.Left == nil {
        return n.Right
    }
    if n.Right == nil {
        return n.Left
    }
    // If the node to be deleted has two children, replace it with the
    // smallest val in the right subtree and delete that node
    smallest := n.Right
    for {
        if smallest.Left == nil {
            break
        }
        smallest = smallest.Left
    }
    n.Val = smallest.Val
    n.Right = n.Right.Delete(smallest.Val)
    return n
}