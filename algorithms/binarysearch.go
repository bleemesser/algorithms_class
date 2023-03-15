package algorithms

func BinarySearch(arr []int, t int) int {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == t {
			return mid
		} else if arr[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}