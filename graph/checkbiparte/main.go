// https://www.geeksforgeeks.org/bipartite-graph/
package main

import "fmt"

// Color encoding
const (
	RED     = -1
	NOCOLOR = 0
	BLUE    = 1
)

func checkBiparte(graphInt map[int][]int, numNodes int) bool {
	var (
		colors  = make([]int, numNodes)
		queue   = make([]int, 0, numNodes)
		visited = make([]bool, numNodes)
	)
	for node := 0; node < numNodes; node++ {
		queue = append(queue, node)
	}

	colors[0] = RED

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] == true {
			continue
		}
		visited[node] = true

		if colors[node] == NOCOLOR {
			colors[node] = RED
		}

		for _, neig := range graphInt[node] {
			if colors[neig] == NOCOLOR {
				// color adjacent node with alternate color
				colors[neig] = -colors[node]
			} else if colors[neig] == colors[node] {
				// check whether adjacent nodes' colors have conflicts
				return false
			}
			queue = append(queue, neig)
		}
	}
	return true
}

func main() {
	var graphBiparte = make(map[int][]int)

	// populate undirected graph with test biparte graph
	graphBiparte[0] = []int{1, 3}
	graphBiparte[1] = []int{0, 2}
	graphBiparte[2] = []int{1, 3}
	graphBiparte[3] = []int{0, 2}
	numNodes := 4

	fmt.Println("Graph Biparte?:", checkBiparte(graphBiparte, numNodes))

	var graphNotBiparte = make(map[int][]int)
	graphNotBiparte[0] = []int{1, 2}
	graphNotBiparte[1] = []int{0, 2}
	graphNotBiparte[2] = []int{0, 1}

	fmt.Println("Graph Biparte?:", checkBiparte(graphNotBiparte, numNodes))
}
