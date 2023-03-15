package algorithms

func QuickSort(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
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
		smaller = QuickSort(smaller)
	}
	if len(larger) > 1 {
		larger = QuickSort(larger)
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
