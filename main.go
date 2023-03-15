package main

import (
	// alg "algorithms_class/algorithms"
	ds "algorithms_class/datastructures"
	"fmt"
	// "time"
)

func main() {
	avltree := ds.AvlTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	// tree := ds.BinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	tree2 := ds.BinaryTree([]int{8,4,2,1,3,6,5,7,16,12,10,9,11,14,13,15,18,17,19,20})
	// avltree.PrintAVLTree(0, 'M')
	// tree.PrintBinaryTree(0, 'M')
	tree2.PrintBinaryTree(0, 'M')
	fmt.Println(avltree.SortToSlice())
	
}
