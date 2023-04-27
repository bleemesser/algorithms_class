package pagerank

import (
	"fmt"
	"math/rand"
	"os"

	// graphviz
	gv "github.com/awalterschulze/gographviz"
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
	sum := 0.0
	for bl := range p.BackLinks {
		sum += pages[p.BackLinks[bl]].Rank / float64(len(pages[p.BackLinks[bl]].Links))
	}
	out := (1 - dampingFactor) + dampingFactor*sum
	return out
}

func (p Page) HasLink(to int) bool {
	for l := range p.Links {
		if p.Links[l] == to {
			return true
		}
	}
	return false
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

	// if we have a super node, add links and backlinks to/from it

	if pr.HasSuper {
		superNode := pr.Pages[0] // super node is always at index 0
		superNode.Links = append(superNode.Links, page.Index)
		superNode.BackLinks = append(superNode.BackLinks, page.Index)
		page.Links = append(page.Links, superNode.Index)
		page.BackLinks = append(page.BackLinks, superNode.Index)
		pr.UpdatePage(0, superNode)
		pr.UpdatePage(page.Index, page)
	}

	return page
}

func (pr *PageRank) AddLink(from, to Page) {
	rfrom := pr.Pages[from.Index]
	rto := pr.Pages[to.Index]

	rfrom.Links = append(rfrom.Links, to.Index)
	rto.BackLinks = append(rto.BackLinks, from.Index)

	pr.UpdatePage(from.Index, rfrom)
	pr.UpdatePage(to.Index, rto)
}

func (pr *PageRank) Run() {
	f, err := os.Create("pagerank.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < pr.Iterations; i++ {
		fmt.Fprintf(f, "Iteration: %d\n", i)
		for i := range pr.Pages {
			pr.Pages[i].Rank = pr.Pages[i].CalculateRank(pr.Pages, pr.DF)
			fmt.Fprintf(f, "%s: Rank: %f, Index: %v, Links: %v, Backlinks: %v\n", pr.Pages[i].Name, pr.Pages[i].Rank, pr.Pages[i].Index, pr.Pages[i].Links, pr.Pages[i].BackLinks)
		}
	}
}

func (pr *PageRank) PrintToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("Average Rank: %f\n\n", pr.CalculateAverageRank()))
	f.WriteString(fmt.Sprintf("Iterations: %d\n\n", pr.Iterations))
	for i := range pr.Pages {
		f.WriteString(fmt.Sprintf("%s: Rank: %f, Index: %v, Links: %v, Backlinks: %v\n", pr.Pages[i].Name, pr.Pages[i].Rank, pr.Pages[i].Index, pr.Pages[i].Links, pr.Pages[i].BackLinks))
	}

	f.WriteString("\n\nSorted by Rank:\n\n")
	sorted := pr.GenerateSortedRankings()
	for i := range sorted {
		f.WriteString(fmt.Sprintf("%s: Rank: %f, Index: %v, Links: %v, Backlinks: %v\n", sorted[i].Name, sorted[i].Rank, sorted[i].Index, sorted[i].Links, sorted[i].BackLinks))
	}

}

func GenerateRandom(n, m int, superNode bool) PageRank { // generate n random pages with m links
	pr := MakePageRank(0.85, 35, 1.0, superNode)

	for i := 0; i < n; i++ {
		index := i
		if superNode {
			index += 1
		}
		pr.AddPage(fmt.Sprintf("Page %v", index))
	}

	for i := 0; i < m; i++ {
		// from shouldn't return 0 if we have a super node
		from := rand.Intn(n)
		to := rand.Intn(n)
		if pr.HasSuper {
			from = rand.Intn(n) + 1
			to = rand.Intn(n) + 1
		}

		// check for self link and duplicate links
		safetyLimit := 1000
		for safetyLimit > 0 && (from == to || pr.Pages[from].HasLink(to)) {
			from = rand.Intn(n)
			to = rand.Intn(n)
			safetyLimit--
		}

		if safetyLimit == 0 {
			continue
		}

		pr.AddLink(pr.Pages[from], pr.Pages[to])

	}

	return pr
}

func (pr PageRank) CalculateAverageRank() float64 {
	sum := 0.0
	for i := range pr.Pages {
		sum += pr.Pages[i].Rank
	}
	return sum / float64(len(pr.Pages))
}

func (pr PageRank) Visualize() {
	// print the graph to console using graphviz
	ast := "digraph G {\n"
	for i := range pr.Pages {
		for j := range pr.Pages[i].Links {
			// add the indices of the links
			ast += fmt.Sprintf("%v -> %v;\n", pr.Pages[i].Index, pr.Pages[i].Links[j])
		}
	}
	ast += "}"

	gvast, err := gv.ParseString(ast)
	if err != nil {
		panic(err)
	}

	graph := gv.NewGraph()
	if err := gv.Analyse(gvast, graph); err != nil {
		panic(err)
	}

	// Render graph to file
	fmt.Printf("graph: %v\n", graph)

}

func QuickSortPages(arr []Page) []Page {
	if len(arr) == 1 || len(arr) == 0 {
		return arr
	}

	middle := arr[len(arr)-1]
	var smaller []Page
	var larger []Page
	var middles []Page
	for _, e := range arr {
		if e.Rank < middle.Rank {
			smaller = append(smaller, e)
		} else if e.Rank > middle.Rank {
			larger = append(larger, e)
		} else if e.Rank == middle.Rank && e.Index != middle.Index {
			middles = append(middles, e)
		}
	}
	if len(smaller) > 1 {
		smaller = QuickSortPages(smaller)
	}
	if len(larger) > 1 {
		larger = QuickSortPages(larger)
	}
	arr = append(smaller, middle)
	arr = append(arr, middles...)

	for i := 0; i < len(larger); i++ {
		arr = append(arr, larger[i])
	}
	return arr
}

func (pr PageRank) GenerateSortedRankings() []Page {
	sorted := QuickSortPages(pr.Pages)
	// reverse the array
	for i := 0; i < len(sorted)/2; i++ {
		sorted[i], sorted[len(sorted)-i-1] = sorted[len(sorted)-i-1], sorted[i]
	}
	return sorted
}
