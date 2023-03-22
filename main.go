package main

import (
	// alg "algorithms_class/algorithms"
	ds "algorithms_class/datastructures"
	"sync"

	// "fmt"
	"github.com/schollz/progressbar"
	"math/rand"
	// "strconv"

	// "strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func timeInsert(size int) time.Duration { // how long does it take to construct a tree of size n
	toInsert := rand.Perm(size)
	tree := ds.BinaryTree([]int{toInsert[0]})
	start := time.Now()
	for i := 1; i < size; i++ {
		tree.Insert(toInsert[i])
	}
	elapsed := time.Since(start)
	return elapsed
}

func timeFindMin(size int) time.Duration { // how long does it take to find the min in a tree of size n
	tree := ds.BinaryTree(rand.Perm(size))
	start := time.Now()
	ds.FindMin(tree)
	elapsed := time.Since(start)
	return elapsed
}

func runTest(iters int, sizes []int) (map[int]time.Duration, map[int]time.Duration) {
    var insertionLock sync.Mutex
    var findMinLock sync.Mutex

    insertionTimes := make(map[int]time.Duration)
    findMinTimes := make(map[int]time.Duration)

    // create a new progress bar
    bar := progressbar.New(len(sizes) * 2 * iters)

    for i := 0; i < iters; i++ {
        var wg sync.WaitGroup

        for _, v := range sizes {
            wg.Add(2)

            go func(v int) {
                defer wg.Done()
				dur := timeInsert(v)
                insertionLock.Lock()
                insertionTimes[v] += dur
				insertionLock.Unlock()
				bar.Add(1) // increment the progress bar

            }(v)

            go func(v int) {
				defer wg.Done()
				dur := timeFindMin(v)
				findMinLock.Lock()
				findMinTimes[v] += dur
				findMinLock.Unlock()
				bar.Add(1) // increment the progress bar
            }(v)
        }

        wg.Wait()
    }

    for k, v := range insertionTimes {
        insertionTimes[k] = v / time.Duration(iters)
    }

    for k, v := range findMinTimes {
        findMinTimes[k] = v / time.Duration(iters)
    }

    return insertionTimes, findMinTimes
}


func main() {
	// array from 100 to 100000 in increments of 1000
	sizes := make([]int, 100)
	for i := 0; i < 100; i++ {
		sizes[i] = (i + 1) * 1000
	}
	iters := 100
	insertionAvg, findminAvg := runTest(iters, sizes)

	// plot
	p := plot.New()
	p.Title.Text = "Binary Tree Insertion vs tree size"
	p.X.Label.Text = "Tree Size"
	p.Y.Label.Text = "Time (microseconds)"

	pts := make(plotter.XYs, len(sizes))

	for i, v := range sizes {
		pts[i].X = float64(v)
		pts[i].Y = float64(insertionAvg[v].Microseconds())
	}

	err := plotutil.AddLinePoints(p, "Insertion", pts)
	if err != nil {
		panic(err)
	}
	// high res plot
	if err := p.Save(16*vg.Inch, 16*vg.Inch, "insertion.png"); err != nil {
		panic(err)
	}

	p2 := plot.New()
	p2.Title.Text = "Binary Tree FindMin vs tree size"
	p2.X.Label.Text = "Tree Size"
	p2.Y.Label.Text = "Time (ns)"

	pts2 := make(plotter.XYs, len(sizes))

	for i, v := range sizes {
		pts2[i].X = float64(v)
		pts2[i].Y = float64(findminAvg[v].Nanoseconds())
	}

	err = plotutil.AddLinePoints(p2, "FindMin", pts2)
	if err != nil {
		panic(err)
	}

	if err := p2.Save(16*vg.Inch, 16*vg.Inch, "findMin.png"); err != nil {
		panic(err)
	}

}
