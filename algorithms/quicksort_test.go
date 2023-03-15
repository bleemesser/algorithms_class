package algorithms

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// test case 1: sort a large array with duplicate elements
	arr := []int{10, 3, 8, 3, 2, 10, 7, 8, 5, 2, 1, 4, 9, 6, 5}
	sortedArr := QuickSort(arr)
	expectedArr := []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 8, 9, 10, 10}
	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expectedArr, sortedArr)
	}

	// test case 2: sort an array with negative elements
	arr = []int{10, -3, 8, -3, 2, 10, -7, 8, 5, 2, 1, 4, -9, 6, 5}
	sortedArr = QuickSort(arr)
	expectedArr = []int{-9, -7, -3, -3, 1, 2, 2, 4, 5, 5, 6, 8, 8, 10, 10}
	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expectedArr, sortedArr)
	}

	// test case 3: sort an array with all identical elements
	arr = []int{5, 5, 5, 5, 5}
	sortedArr = QuickSort(arr)
	expectedArr = []int{5, 5, 5, 5, 5}
	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Errorf("Test case 3 failed: expected %v, got %v", expectedArr, sortedArr)
	}

	// test case 4: sort an array with 0 elements
	arr = []int{}
	sortedArr = QuickSort(arr)
	expectedArr = []int{}
	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Errorf("Test case 4 failed: expected %v, got %v", expectedArr, sortedArr)
	}

	// test case 5: sort an array with 1 element
	arr = []int{5}
	sortedArr = QuickSort(arr)
	expectedArr = []int{5}
	if !reflect.DeepEqual(sortedArr, expectedArr) {
		t.Errorf("Test case 5 failed: expected %v, got %v", expectedArr, sortedArr)
	}

}