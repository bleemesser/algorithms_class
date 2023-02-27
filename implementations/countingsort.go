package main

import (
	"fmt"
	"time"
)

func countingSort(arr []int) []int {
	var max int
	var min int = 1 << 32
	for _, e := range arr {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	buckets := make([]int,max-min+1)
	for _, e := range arr {
		buckets[e-min]++
	}

	var out []int
	for i, e := range buckets {
		if e != 0 {
			for j := 0; j < e; j++ {
				out = append(out, i+min)
			}
		}
	}
	return out
}

func main() {
	randarr := []int{9,3,6,2,8,7,3,4}
	start := time.Now()
	countingSort(randarr)
	fmt.Println(time.Since(start))
}