package algorithms

func InsertionSort(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// InsertionSort works through what I would call cascading swaps. It starts by comparing the first 
// two elements and swapping them if necessary. Then, since the first two elements 
// make up a sorted list, it compares the third element with the second element.
// If the third element is smaller than the second element, it swaps them. Then,
// it compares the third element with the first element and swaps them if necessary.
// For all following elements, it compares them with the elements to their left
// and moves them to the left until they are in the correct position.

// The fastest input for InsertionSort is an already sorted list. In this case,
// since it will only have to compare each element with the one to its left once,
// it will run in O(n) time. The slowest input is a reverse sorted list. In this case,
// it will have to compare each element with every other element before it, so it will
// run in O(n^2) time.

// Other attributes: stable & in-place. It is also adaptive, but I'm not sure how best to 
// determine that property.
