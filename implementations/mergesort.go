package main

import (
	"fmt"
	"math"
	"time"
)

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := math.Floor(float64(len(arr) / 2))
	left := mergeSort(arr[:int(mid)])
	right := mergeSort(arr[int(mid):])
	li := 0
	ri := 0
	var result []int
	for li < len(left) && ri < len(right) {
		if left[li] <= right[ri] {
			result = append(result, left[li])
			li += 1
		} else {
			result = append(result, right[ri])
			ri += 1
		}
	}
	for _,e := range left[int(li):] {
		result = append(result, e)
	}
	for _,e := range right[int(ri):] {
		result = append(result, e)
	}
	return result
}
func main() {
	randarr := []int{9,3,6,2,8,7,3,4}
	start := time.Now()
	mergeSort(randarr)
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}