package tarjan

type Graph struct {
	adj [][]int
}

func NewGraph(numNodes int) *Graph {
	return &Graph{adj: make([][]int, numNodes)}
}

// AddEdge from "from" to "to" node. Directed edge, 0-indexed
func (g *Graph) AddEdge(from, to int) {
	g.adj[from] = append(g.adj[from], to)
}

// SCC returns strongly connected components
func (g *Graph) SCC() [][]int {
	n := len(g.adj)
	scc := &scc{
		adj:       g.adj,
		anc:       make([]int, n),
		ord:       make([]int, n),
		isInStack: make([]bool, n),
		time:      0,
		stack:     []int{},
		result:    [][]int{},
	}
	for i := 0; i < n; i++ {
		if scc.ord[i] == 0 {
			scc.dfs(i)
		}
	}
	return scc.result
}

type scc struct {
	adj       [][]int
	anc       []int   // highest ancestor in the current dfs tree
	ord       []int   // order in which node was discovered
	isInStack []bool  // for dfs traversal to indicate whether nodes are within same dfs tree
	time      int     // to indicate time at which node was discovered
	stack     []int   // keep vertices traversed for current dfs traversal
	result    [][]int // stores strongly connected components, each entry stores node list which belong to one strongly connected component
}

func (s *scc) dfs(from int) {
	s.time++
	s.ord[from] = s.time
	s.anc[from] = s.time
	s.isInStack[from] = true
	s.stack = append(s.stack, from)

	for _, neigh := range s.adj[from] {
		if s.ord[neigh] == 0 { // not visited yet
			s.dfs(neigh)
			s.anc[from] = min(s.anc[from], s.anc[neigh])
		} else if s.isInStack[neigh] { // if we visited in the same dfs traversal before point to it as highest ancestor
			s.anc[from] = min(s.anc[from], s.ord[neigh])
		}
	}
	if s.ord[from] == s.anc[from] { // means vertex is root of scc (it could be alone there)
		component := []int{}
		for len(s.stack) != 0 { // all scc under this tree are already popped, so all stack elements inserted after "from" form scc and their anc value equal from
			lastIndex := len(s.stack) - 1
			node := s.stack[lastIndex]
			s.stack = s.stack[:lastIndex] //pop from stack
			s.isInStack[node] = false
			component = append(component, node)
			if node == from {
				break
			}
		}
		s.result = append(s.result, component)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// once we construct sccs, we will use dagEdges to represent connections between sccs (cross edges)
type dagEdge struct {
	from int
	to   int
}
