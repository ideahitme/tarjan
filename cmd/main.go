package main

import (
	"fmt"

	"github.com/ideahitme/tarjan"
)

func main() {
	g := tarjan.NewGraph(6)
	g.AddEdge(1, 0)
	g.AddEdge(0, 2)
	g.AddEdge(2, 1)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(0, 4)
	g.AddEdge(5, 4)
	fmt.Println(g.SCC())
}
