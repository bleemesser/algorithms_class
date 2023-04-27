package main

import (
	pg "algorithms_class/pagerank"
)

func main() {
	// pr := pg.MakePageRank(0.85, 35, 1.0, false)	// this is the example from the class slides
	// pageA := pr.AddPage("A")
	// pageB := pr.AddPage("B")
	// pageC := pr.AddPage("C")
	// pageD := pr.AddPage("D")

	// pr.AddLink(pageA, pageB)
	// pr.AddLink(pageA, pageC)
	// pr.AddLink(pageB, pageC)
	// pr.AddLink(pageC, pageA)
	// pr.AddLink(pageD, pageC)

	// generate a random configuration of pages and links
	pr := pg.GenerateRandom(100, 100, true) // random generation of 100 pages with 100 links in total

	pr.Run()

	// pr.Visualize()

	pr.PrintToFile("pagerank.txt")


	// SEE OUTPUT FILE, results are correct to the example 
	// if you uncomment the above code (and comment out the random generation)

}