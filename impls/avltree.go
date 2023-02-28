package impls
import (
	"fmt"
 	"io"
)
type BinaryTreeNode struct {
	Val, Height   int
	Left, Right  *BinaryTreeNode
}

func (n *BinaryTreeNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *BinaryTreeNode) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
	// if return is less than -1 or greater than 1, tree is unbalanced
	// negative meas right is heavier, positive means left is heavier
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Greatest(n *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n
	}
	return Greatest(n.Right)
}

func CreateNode(val int) *BinaryTreeNode {
	return &BinaryTreeNode{
		Val:    val,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}
}

func (n *BinaryTreeNode) UpdateHeight() {
	if n == nil {
		return
	}
	n.Height = 1 + Max(n.Left.GetHeight(), n.Right.GetHeight())
}

func RightRotate(n *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	left := n.Left
	right := left.Right
	left.Right = n
	n.Left = right

	n.UpdateHeight()
	left.UpdateHeight()
	return left
}

func LeftRotate(n *BinaryTreeNode) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	right := n.Right
	left := right.Left
	right.Left = n
	n.Right = left

	n.UpdateHeight()
	right.UpdateHeight()
	return right
}

func RotateInsert(n *BinaryTreeNode, val int) *BinaryTreeNode {
	if n == nil {
		return CreateNode(val)
	}
	n.UpdateHeight()
	weight := n.BalanceFactor()
	if weight > 1 && val < n.Left.Val {
		return RightRotate(n)
	}
	if weight < -1 && val > n.Right.Val {
		return LeftRotate(n)
	}
	if weight > 1 && val > n.Left.Val {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}
	if weight < -1 && val < n.Right.Val {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}
	return n
}

func RotateDelete(n *BinaryTreeNode) *BinaryTreeNode {
	n.UpdateHeight()
	weight := n.BalanceFactor()

	if weight > 1 && n.Left.BalanceFactor() >= 0 {
		return RightRotate(n)
	}
	if weight > 1 && n.Left.BalanceFactor() < 0 {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}
	if weight < -1 && n.Right.BalanceFactor() <= 0 {
		return LeftRotate(n)
	}
	if weight < -1 && n.Right.BalanceFactor() > 0 {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}
	return n
}

func InsertNode(n *BinaryTreeNode, val int) *BinaryTreeNode {
	if n == nil {
		return CreateNode(val)
	}
	if n.Val == val {
		return nil
	}
	if val > n.Val {
		n.Right = InsertNode(n.Right, val)
	}
	if val < n.Val {
		n.Left = InsertNode(n.Left, val)
	}
	return RotateInsert(n, val)
}

func DeleteNode(n *BinaryTreeNode, val int) *BinaryTreeNode {
	if n == nil {
		return nil
	}
	if val > n.Val {
		n.Right = DeleteNode(n.Right, val)
	} else if val < n.Val {
		n.Left = DeleteNode(n.Left, val)
	} else {
		if n.Left != nil || n.Right != nil {
			next := Greatest(n.Left)
			val := next.Val
			n.Left = DeleteNode(n.Left, val)
			n.Val = val
		} else if n.Left != nil || n.Right != nil {
			if n.Left != nil {
				n = n.Left
			} else {
				n = n.Right
			}
		} else if n.Left == nil && n.Right == nil {
			n = nil
		}
	}
	if n == nil {
		return nil
	}
	return RotateDelete(n)
}

func AvlTree(arr []int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, val := range arr {
		root = InsertNode(root, val)
	}
	return root
}

func PrintTree(w io.Writer, node *BinaryTreeNode, ns int, ch rune) {
    if node == nil {
        return
    }
 
    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
    PrintTree(w, node.Left, ns+2, 'L')
    PrintTree(w, node.Right, ns+2, 'R')
}

