package pagerank

import (
	"fmt"
	"os"
)

// PR(A) = (1-d) + d(PR(T1)/C(T1) + ... + PR(Tn)/C(Tn))
// PR() is the PageRank
// A is the page
// C() is the number of links on the page (outgoing links)
// T1...Tn are the pages that link to A

type Page struct {
	Name      string
	Index     int
	Links     []int
	BackLinks []int
	Rank      float64
}

func NewPage(name string) Page {
	return Page{Name: name, Links: []int{}, BackLinks: []int{}, Rank: 0.0}
}

func (p *Page) CalculateRank(pages []Page, dampingFactor float64) float64 {
	fmt.Println("Page ", p.Name, " has rank ", p.Rank, " and ", len(p.BackLinks), " backlinks")
	sum := 0.0
	for bl := range p.BackLinks {
		fmt.Println("Calculating rank using backlink ", pages[p.BackLinks[bl]].Name, " with rank ", pages[p.BackLinks[bl]].Rank, " and ", len(pages[p.BackLinks[bl]].Links), " links")
		sum += pages[p.BackLinks[bl]].Rank / float64(len(pages[p.BackLinks[bl]].Links))
	}
	out := (1 - dampingFactor) + dampingFactor*sum
	fmt.Println("Page ", p.Name, " has new rank ", out)
	return out
}

type PageRank struct { // this is organized so only the pagerank struct can modify its data
	DF         float64
	Pages      []Page
	Iterations int
	InitRank   float64
	HasSuper   bool
}

func (pr *PageRank) UpdatePage(index int, page Page) {
	pr.Pages[index] = page
}

func MakePageRank(dampingFactor float64, iterations int, initialRank float64, addSuperNode bool) PageRank {
	pr := PageRank{DF: dampingFactor, Pages: []Page{}, Iterations: iterations, InitRank: initialRank, HasSuper: addSuperNode}
	if addSuperNode {
		pr.AddSuperNode()
	}
	return pr
}

func (pr *PageRank) AddSuperNode() {
	page := NewPage("SuperNode")
	page.Index = 0
	page.Rank = 1.0
	page.Links = []int{}
	page.BackLinks = []int{}
	pr.Pages = append(pr.Pages, page)
}

func (pr *PageRank) AddPage(name string) Page {
	page := NewPage(name)
	page.Index = len(pr.Pages)
	page.Rank = pr.InitRank
	page.Links = []int{}
	page.BackLinks = []int{}
	pr.Pages = append(pr.Pages, page)

	// if we have a super node, add a link from it to this page

	if pr.HasSuper {
		superNode := pr.Pages[0] // super node is always at index 0
		superNode.Links = append(superNode.Links, page.Index)
		page.BackLinks = append(page.BackLinks, superNode.Index)
		pr.UpdatePage(0, superNode)
		pr.UpdatePage(page.Index, page)
	}

	return page
}

func (pr *PageRank) AddLink(from, to Page) {
	// fmt.Println("Adding link from ", from.Name, " to ", to.Name)
	// fmt.Println("from index: ", from.Index, " to index: ", to.Index)

	rfrom := pr.Pages[from.Index]
	rto := pr.Pages[to.Index]

	rfrom.Links = append(rfrom.Links, to.Index)
	rto.BackLinks = append(rto.BackLinks, from.Index)
	
	pr.UpdatePage(from.Index, rfrom)
	pr.UpdatePage(to.Index, rto)

	// fmt.Println("Finished adding link from ", from.Name, " to ", to.Name)
	// fmt.Println("FROM LINKS: ", pr.Pages[from.Index].Links)
	// fmt.Println("TO BACKLINKS: ", pr.Pages[to.Index].BackLinks)
}

func (pr *PageRank) Run() {
	f, err := os.Create("pagerank.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < pr.Iterations; i++ {
		fmt.Fprintf(f, fmt.Sprintf("Iteration: %d\n", i))
		for i := range pr.Pages {
			pr.Pages[i].Rank = pr.Pages[i].CalculateRank(pr.Pages, pr.DF)
			fmt.Fprintf(f, fmt.Sprintf("%s: Rank: %f, Index: %v, Links: %v, Backlinks: %v\n", pr.Pages[i].Name, pr.Pages[i].Rank, pr.Pages[i].Index, pr.Pages[i].Links, pr.Pages[i].BackLinks))
		}
		fmt.Println("Iteration: ", i)
	}
}

// print to file
func (pr *PageRank) PrintToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for i := range pr.Pages {
		f.WriteString(fmt.Sprintf("After %v iterations:\n%s: Rank: %f, Index: %v, Links: %v, Backlinks: %v\n", pr.Iterations, pr.Pages[i].Name, pr.Pages[i].Rank, pr.Pages[i].Index, pr.Pages[i].Links, pr.Pages[i].BackLinks))
	}
}
