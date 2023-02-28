package impls
import (
	"fmt"
	"os"
)

// Package to implement AVL tree

type Node struct {
	Val, height int
	Left, Right *Node
}
// get height of node
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}
// get balance factor of node: negative if right is heavier, positive if left is heavier
// if return is less than -1 or greater than 1, tree is unbalanced
func (n *Node) balanceFactor() int {
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
	return &Node {
		Val: val,
		height: 1,
		Left: nil,
		Right: nil,
	}
}
// balance after insertion, given the newly inserted value
func (n *Node) iBalance(val int) *Node {
	if n == nil {
		return nil
	}
	n.updateHeight()
	weight := n.balanceFactor()
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
	weight := n.balanceFactor()
	if weight > 1 && n.Left.balanceFactor() >= 0 {
		return rRotate(n)
	}
	if weight < -1 && n.Right.balanceFactor() <= 0 {
		return lRotate(n)
	}
	if weight > 1 && n.Left.balanceFactor() < 0 {
		// left right rotation
		n.Left = lRotate(n.Left)
		return rRotate(n)
	}
	if weight < -1 && n.Right.balanceFactor() > 0 {
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
func (n *Node) PrintTree(ns int, ch rune) {
	if n == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(os.Stdout, " ")
	}
	fmt.Fprintf(os.Stdout, "%c:%v\n", ch, n.Val)
	n.Left.PrintTree(ns+2, 'L')
	n.Right.PrintTree(ns+2, 'R')
}
// create a new tree from a list of values. This is the primary function to be used
func AvlTree(vals []int) *Node {
	var root *Node
	for _, val := range vals {
		root = root.InsertNode(val)
	}
	return root
}
// get a list of values in the tree in order lowest to highest
func (n *Node) ExportToSlice() []int {
	if n == nil {
		return nil
	}
	var vals []int
	vals = append(vals, n.Left.ExportToSlice()...)
	vals = append(vals, n.Val)
	vals = append(vals, n.Right.ExportToSlice()...)
	return vals
}


