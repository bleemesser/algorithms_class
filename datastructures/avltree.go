package datastructures

import (
	"fmt"
	"os"
)

// Package to implement AVL tree

type Node struct {
	Val, height int
	Left, Right *Node
}

// return height of referenced node, needed for getting height of the left and right subtrees @ their pointers
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// get weight of node: negative if right is heavier, positive if left is heavier
// if return is less than -1 or greater than 1, tree is unbalanced
func (n *Node) weight() int {
	if n == nil {
		return 0
	}
	return n.Left.Height() - n.Right.Height()
}

// get max of two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// update height of node, for use after rotations
func (n *Node) updateHeight() {
	if n == nil {
		return
	}
	n.height = 1 + max(n.Left.Height(), n.Right.Height())
}

// defining roations
// right rotation
func rRotate(n *Node) *Node {
	if n == nil {
		return nil
	}
	left := n.Left
	right := left.Right
	left.Right = n
	n.Left = right
	n.updateHeight()
	left.updateHeight()
	return left
}

// left rotation
func lRotate(n *Node) *Node {
	if n == nil {
		return nil
	}
	right := n.Right
	left := right.Left
	right.Left = n
	n.Right = left
	n.updateHeight()
	right.updateHeight()
	return right
}

// create a new node given a value
func createNode(val int) *Node {
	return &Node{
		Val:    val,
		height: 1,
		Left:   nil,
		Right:  nil,
	}
}

// balance after insertion, given the newly inserted value
func (n *Node) iBalance(val int) *Node {
	if n == nil {
		return nil
	}
	n.updateHeight()
	weight := n.weight()
	if weight > 1 && val < n.Left.Val {
		return rRotate(n)
	}
	if weight < -1 && val > n.Right.Val {
		return lRotate(n)
	}
	if weight > 1 && val > n.Left.Val {
		// left right rotation
		n.Left = lRotate(n.Left)
		return rRotate(n)
	}
	if weight < -1 && val < n.Right.Val {
		// right left rotation
		n.Right = rRotate(n.Right)
		return lRotate(n)
	}
	return n
}

// insert a node into the tree
func (n *Node) InsertNode(val int) *Node {
	if n == nil {
		return createNode(val)
	}
	if val < n.Val {
		n.Left = n.Left.InsertNode(val)
	} else if val > n.Val {
		n.Right = n.Right.InsertNode(val)
	} else if val == n.Val {
		return n.iBalance(val)
	}
	return n.iBalance(val)
}

// balance after deletion
func (n *Node) dBalance() *Node {
	if n == nil {
		return nil
	}
	n.updateHeight()
	weight := n.weight()
	if weight > 1 && n.Left.weight() >= 0 {
		return rRotate(n)
	}
	if weight < -1 && n.Right.weight() <= 0 {
		return lRotate(n)
	}
	if weight > 1 && n.Left.weight() < 0 {
		// left right rotation
		n.Left = lRotate(n.Left)
		return rRotate(n)
	}
	if weight < -1 && n.Right.weight() > 0 {
		// right left rotation
		n.Right = rRotate(n.Right)
		return lRotate(n)
	}
	return n
}

// recrsively get node with the greatest value in the tree
func greatest(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n
	}
	return greatest(n.Right)
}

// delete a node by value from the tree
func (n *Node) DeleteNode(val int) *Node {
	if n == nil {
		return nil
	}
	if val > n.Val {
		n.Right = n.Right.DeleteNode(val)
	} else if val < n.Val {
		n.Left = n.Left.DeleteNode(val)
	} else {
		if n.Left != nil || n.Right != nil {
			next := greatest(n.Left)
			v := next.Val
			n.Left = n.Left.DeleteNode(v)
			n.Val = v
		} else if n.Left == nil && n.Right == nil {
			return nil
		}
	}
	return n.dBalance()
}

// pretty-print the tree. ns is the number of spaces to indent, ch is the character to print
// should be called with 0, 'M' always, as the root node is the middle of the tree and isn't indented
func (n *Node) PrintAVLTree(ns int, ch rune) {
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
	n.Left.PrintAVLTree(ns+2, 'L')
	n.Right.PrintAVLTree(ns+2, 'R')
}

// create a new tree from a list of values. This is the primary function to be used
func AvlTree(vals []int) *Node {
	var root *Node
	for _, val := range vals {
		root = root.InsertNode(val)
	}
	return root
}

// export the tree to a slice of values, sorted low to high
func (n *Node) SortToSlice() []int {
	if n == nil {
		// fmt.Println("CALLED NIL")
		return nil
	}
	// fmt.Println("CALLED ", n.Val)

	var vals []int
	// n.PrintAVLTree(0,'M')
	vals = append(vals, n.Left.SortToSlice()...)
	// fmt.Println("Appended left ", vals)
	vals = append(vals, n.Val)
	// fmt.Println("Appended middle ",vals)
	vals = append(vals, n.Right.SortToSlice()...)
	// fmt.Println("Appended right ", vals)
	// fmt.Println("Finished ", n.Val)
	return vals
}

func FindMax(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Right == nil {
		return root.Val
	}
	return FindMax(root.Right)
}

func FindMin(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return root.Val
	}
	return FindMin(root.Left)
}

func FindVal(root *Node, val int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	if val < root.Val {
		return FindVal(root.Left, val)
	}
	return FindVal(root.Right, val)
}