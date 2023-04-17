package pagerank

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

func (p *Page) AddLink(to int) {
	p.Links = append(p.Links, to)
}

func (p *Page) AddBackLink(from int) {
	p.BackLinks = append(p.BackLinks, from)
}

func (p *Page) CalculateRank(pages []Page, dampingFactor float64) float64 {
	rank := (1.0 - dampingFactor)
	for _, i := range p.BackLinks {
		rank += dampingFactor * (pages[i].Rank / float64(len(pages[i].Links)))
	}
	return rank
}

type PageRank struct {
	DF         float64
	Pages      []Page
	Iterations int
}

func (pr *PageRank) AddPage(name string) int {
	p := Page{Name: name, Index: len(pr.Pages), Rank: 1.0}
	pr.Pages = append(pr.Pages, p)
	return p.Index
}

func (pr *PageRank) AddLink(from, to int) {
	pr.Pages[from].AddLink(to)
	pr.Pages[to].AddBackLink(from)
}

func MakePageRank(dampingFactor float64, iterations int, pages []Page, initialRank float64, addSuperNode bool) PageRank {
	pr := PageRank{}
	pr.DF = dampingFactor
	pr.Iterations = iterations

	pr.Pages = pages
	
	for i := range pr.Pages {
		pr.Pages[i].Rank = initialRank
	}

	if addSuperNode {
		superNode := pr.AddPage("SuperNode")
		for i := range pr.Pages {
			pr.AddLink(superNode, i)
			pr.AddLink(i, superNode)
		}
	}

	return pr
}

