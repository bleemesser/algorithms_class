package impls

func CountingSort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
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
	// fmt.Println(max, min)
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
