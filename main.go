package main
import (
	"algorithms_class/impls"
	"fmt"

)

func main() {
	tree := impls.AvlTree([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	fmt.Println(tree.ExportToSlice())
	
}