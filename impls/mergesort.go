package impls

func MergeSort(arr []int) []int {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}
	mid := len(arr) >> 1
	left := MergeSort(arr[:int(mid)])
	right := MergeSort(arr[int(mid):])
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
	result = append(result, left[int(li):]...)
	result = append(result, right[int(ri):]...)
	return result
}