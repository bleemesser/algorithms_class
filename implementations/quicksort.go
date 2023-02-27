package main

import (
	"fmt"
	"time"
	// "math/rand"
)

func quickSort(arr []int) []int {
	middle := arr[len(arr)-1]
	var smaller []int
	var larger []int
	numMiddles := 0
	for _, e := range arr {
		if e < middle {
			smaller = append(smaller, e)
		} else if e > middle {
			larger = append(larger, e)
		} else if e == middle {
			numMiddles += 1
		}
	}
	if len(smaller) > 1 {
		smaller = quickSort(smaller)
	}
	if len(larger) > 1 {
		larger = quickSort(larger)
	}
	arr = append(smaller, middle)

	for i := 0; i < numMiddles-1; i++ {
		arr = append(arr, middle)
	}
	for i := 0; i < len(larger); i++ {
		arr = append(arr, larger[i])
	}
	return arr
}

func main() {
	randarr := []int{9,3,6,2,8,7,3,4}
	start := time.Now()
	quickSort(randarr)
	fmt.Println(time.Since(start))
}
