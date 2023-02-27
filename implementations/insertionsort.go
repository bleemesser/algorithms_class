package main

import (
	"fmt"
	// "time"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
func main() {
	fmt.Println(insertionSort([]int{}))
}
// func main() {
// 	// time calling quicksort on the array
// 	// randarr := rand.Perm(100)
// 	randarr := []int{9,3,6,2,8,7,3,4}
// 	start := time.Now()
// 	insertionSort(randarr)
// 	elapsed := time.Since(start)
// 	// fmt.Println("Array: ", randarr, "\nSorted Array: ", arr, "\nTime elapsed: ", elapsed)
// 	fmt.Println("Time elapsed: ", elapsed)
// }
