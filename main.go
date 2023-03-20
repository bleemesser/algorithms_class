package main

import (
	// alg "algorithms_class/algorithms"
	ds "algorithms_class/datastructures"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"strconv"
)

func timeInsert(size int) time.Duration{
	var totalSum time.Duration
	for  i := 0; i < size; i++ {
		toInsert := rand.Perm(size)
		tree := ds.BinaryTree([]int{toInsert[0]})
		var insertTime []time.Duration
		for i := 0; i < size; i++ {
			start := time.Now()
			tree.Insert(toInsert[i])
			elapsed := time.Since(start)
			insertTime = append(insertTime, elapsed)
		}
		for _, v := range insertTime {
			totalSum += v
		}
	}
	avg := totalSum / time.Duration(size)
	return avg
}

func timeFindMin(size int) time.Duration{
	var totalSum time.Duration
	for  i := 0; i < size; i++ {
		toInsert := rand.Perm(size)
		tree := ds.BinaryTree([]int{toInsert[0]})
		var findMinTime []time.Duration
		for i := 0; i < size; i++ {
			start := time.Now()
			ds.FindMin(tree)
			elapsed := time.Since(start)
			findMinTime = append(findMinTime, elapsed)
		}
		for _, v := range findMinTime {
			totalSum += v
		}
	}
	avg := totalSum / time.Duration(size)
	return avg
}

func timeFindVal(size int) time.Duration{
	var totalSum time.Duration
	for  i := 0; i < size; i++ {
		toInsert := rand.Perm(size)
		tree := ds.BinaryTree([]int{toInsert[0]})
		var findValTime []time.Duration
		for i := 0; i < size; i++ {
			start := time.Now()
			ds.FindVal(tree, toInsert[i])
			elapsed := time.Since(start)
			findValTime = append(findValTime, elapsed)
		}
		for _, v := range findValTime {
			totalSum += v
		}
	}
	
	return totalSum
}



func main() {
	overallAvg := make(map[string]time.Duration)

	iters := 1

	for j := 0; j < iters; j++ {
		//BINARY TREE
		sum := timeInsert(10000)
		overallAvg["Binary Tree Insertion 10000"] += sum
		sum = timeInsert(20000)
		overallAvg["Binary Tree Insertion 20000"] += sum
		sum = timeInsert(40000)
		overallAvg["Binary Tree Insertion 40000"] += sum
		sum = timeInsert(80000)
		overallAvg["Binary Tree Insertion 80000"] += sum


		sum = timeFindMin(10000)
		overallAvg["Binary Tree Find Min 10000"] += sum
		sum = timeFindMin(20000)
		overallAvg["Binary Tree Find Min 20000"] += sum
		sum = timeFindMin(40000)
		overallAvg["Binary Tree Find Min 40000"] += sum
		sum = timeFindMin(80000)
		overallAvg["Binary Tree Find Min 80000"] += sum


		sum = timeFindVal(10000)
		overallAvg["Binary Tree Find Val 10000"] += sum
		sum = timeFindVal(20000)
		overallAvg["Binary Tree Find Val 20000"] += sum
		sum = timeFindVal(40000)
		overallAvg["Binary Tree Find Val 40000"] += sum
		sum = timeFindVal(80000)
		overallAvg["Binary Tree Find Val 80000"] += sum
	}

	for k, v := range overallAvg {
		num,_ := strconv.Atoi(strings.Split(k, " ")[3])
		totalIters := iters * num
		fmt.Println(time.Duration(totalIters))
		overallAvg[k] = v / time.Duration(totalIters)
	}
	

	fmt.Println("Overall averages: ")
	for k, v := range overallAvg {
		fmt.Println(k, v)
	}
	

}
