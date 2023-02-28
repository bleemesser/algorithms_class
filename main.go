package main
import (
	"algorithms_class/impls"
	// "fmt"
	"os"

)

func main() {
	impls.PrintTree(os.Stdout, impls.AvlTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 0, 'M')
}