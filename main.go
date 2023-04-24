package main

import (
	pg "algorithms_class/pagerank"
)

func main() {
	pr := pg.MakePageRank(0.85, 35, 0.25, false)	
	pageA := pr.AddPage("A")
	pageB := pr.AddPage("B")
	pageC := pr.AddPage("C")
	pageD := pr.AddPage("D")

	pr.AddLink(pageA, pageB)
	pr.AddLink(pageA, pageC)
	pr.AddLink(pageB, pageC)
	pr.AddLink(pageC, pageA)
	pr.AddLink(pageD, pageC)

	pr.Run()

	pr.PrintToFile("pagerank.txt")


	// see output file, results are correct to the example.

}