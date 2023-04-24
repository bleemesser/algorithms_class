// package main

// import (
// 	// alg "algorithms_class/algorithms"
// 	ds "algorithms_class/datastructures"
// 	"sync"

// 	"fmt"
// 	"math/rand"

// 	"github.com/schollz/progressbar"

// 	// "strconv"

// 	// "strings"
// 	"time"

// 	"gonum.org/v1/plot"
// 	"gonum.org/v1/plot/plotter"
// 	"gonum.org/v1/plot/plotutil"
// 	"gonum.org/v1/plot/vg"
// )

// func runTest(iters int, findMinIters int, insertionSizes []int, findMinSizes []int) (map[int]time.Duration, map[int]time.Duration) {
// 	var constructionLock sync.Mutex
// 	insertionTimes := make(map[int]time.Duration)
// 	findMinTimes := make(map[int]time.Duration)

// 	// create a new progress bar for the trial loop
// 	bar := progressbar.New(iters * len(insertionSizes))

// 	// create a progress bar for constructing the trees
// 	bar2 := progressbar.New(len(insertionSizes))

// 	// create a progress bar for finding the min
// 	bar3 := progressbar.New(len(findMinSizes) * findMinIters)
// 	// create random trees for each size, use multiple threads
// 	trees := make(map[int]*ds.Node)
// 	var cg sync.WaitGroup
// 	for _, v := range insertionSizes {
// 		cg.Add(1)
// 		go func(v int) {
// 			defer cg.Done()
// 			defer bar2.Add(1)
// 			tree := ds.AvlTree(rand.Perm(v))
// 			constructionLock.Lock()
// 			trees[v] = tree
// 			constructionLock.Unlock()
// 		}(v)

// 	}
// 	cg.Wait()

// 	bar.Reset()
// 	fmt.Println(trees)
// 	for i := 0; i < iters; i++ {
// 		for size, v := range trees {
// 			// insert
// 			toInsert := rand.Intn(size)
// 			start := time.Now()
// 			v.InsertNode(toInsert)
// 			elapsed := time.Since(start)
// 			insertionTimes[size] += elapsed
// 			v.DeleteNode(toInsert)
// 			bar.Add(1)
// 		}
// 	}

// 	// find min
// 	for i := 0; i < findMinIters; i++ {
// 		for _, size := range findMinSizes {
// 			tree := ds.AvlTree(rand.Perm(size))
// 			start := time.Now()
// 			ds.FindAVLMin(tree)
// 			elapsed := time.Since(start)
// 			findMinTimes[size] += elapsed
// 			bar3.Add(1)
// 		}
// 	}

// 	for k, v := range insertionTimes {
// 		insertionTimes[k] = v / time.Duration(iters)
// 	}

// 	for k, v := range findMinTimes {
// 		findMinTimes[k] = v / time.Duration(findMinIters)
// 	}

// 	return insertionTimes, findMinTimes
// }

// func main() {
// 	// array from 1000 to 40 million in doubling incrememts 1000, 2000, 4000, etc
// 	insertionSizes := make([]int, 0)
// 	for i := 1000; i <= 40000000; i *= 2 {
// 		insertionSizes = append(insertionSizes, i)
// 	}
// 	// array from 1000 to 100,000 in increments of 1000
// 	findMinSizes := make([]int, 0)
// 	for i := 100000; i <= 1000000; i += 100000 {
// 		findMinSizes = append(findMinSizes, i)
// 	}
// 	fmt.Println(insertionSizes)
// 	fmt.Println(findMinSizes)
// 	iters := 2000000
// 	findMinIters := 100
// 	insertionAvg, findminAvg := runTest(iters, findMinIters, insertionSizes, findMinSizes)

// 	// plot
// 	p := plot.New()
// 	p.Title.Text = "AVL Tree Insertion time vs tree size"
// 	p.X.Label.Text = "Tree Size"
// 	p.Y.Label.Text = "Time (ns)"

// 	pts := make(plotter.XYs, len(insertionSizes))

// 	for i, v := range insertionSizes {
// 		pts[i].X = float64(v)
// 		pts[i].Y = float64(insertionAvg[v].Nanoseconds())
// 	}

// 	err := plotutil.AddLinePoints(p, "Insertion", pts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// high res plot
// 	if err := p.Save(16*vg.Inch, 16*vg.Inch, "insertion.png"); err != nil {
// 		panic(err)
// 	}

// 	p2 := plot.New()
// 	p2.Title.Text = "AVL Tree FindMin vs tree size"
// 	p2.X.Label.Text = "Tree Size"
// 	p2.Y.Label.Text = "Time (ns)"

// 	pts2 := make(plotter.XYs, len(findMinSizes))

// 	for i, v := range findMinSizes {
// 		pts2[i].X = float64(v)
// 		pts2[i].Y = float64(findminAvg[v].Nanoseconds())
// 	}

// 	err = plotutil.AddLinePoints(p2, "FindMin", pts2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := p2.Save(16*vg.Inch, 16*vg.Inch, "findMin.png"); err != nil {
// 		panic(err)
// 	}

// }
