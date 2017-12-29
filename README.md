### Tarjan algorithm

Implementation of Tarjan algorithm in Go. The purpose of the algorithm is to produce all SCCs (strong connected components) in a directed graph. Complexity of the algorithm `O(V+E)` 

#### How to use it: 

Get all SCCs:

```go
package main

import (
	"github.com/ideahitme/tarjan"
	"fmt"
)

func main() {
	g := tarjan.NewGraph(5)
	g.AddEdge(1, 0)
	g.AddEdge(0, 2)
	g.AddEdge(2, 1)
	g.AddEdge(0, 3)
	g.AddEdge(3, 4)
	fmt.Println(g.SCC())
}
```

prints 

```bash
[[4] [3] [1 2 0]]
```

Means vertices `1,2,0` form strongly connected component and `4` and `3` are vertices which form strongly connected components of size 1.  