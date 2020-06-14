// https://www.geeksforgeeks.org/maximum-number-edges-added-tree-stays-bipartite-graph/
package main

import "fmt"

// Color encoding
const (
	RED     = -1
	NOCOLOR = 0
	BLUE    = 1
)

func maxAdditionalEdgeBiparte(edges [][2]int, numNodes int) int {
	colors := make([]int, numNodes)
	for _, edge := range edges {
		u, v := edge[0], edge[1]

		if colors[u] == NOCOLOR && colors[v] == NOCOLOR {
			// case when both nodes are uncolored
			colors[u], colors[v] = RED, BLUE
		} else if (colors[v] != NOCOLOR) && (colors[v] == colors[u]) {
			// case when the input edges itself create color conflict
			// such that color doesn't alternate
			fmt.Printf("Invalid edge from node (%d) to node (%d). Cannot make biparte\n", u, v)
			return 0
		} else if colors[u] != NOCOLOR {
			// if one of them is colored, color the other with
			// opposite color
			colors[v] = -colors[u]
		} else if colors[v] != NOCOLOR {
			// same as above
			colors[u] = -colors[v]
		}
	}

	// Formula for maximum edges for biparte: maxEdges = B x R
	// assuming all nodes has been assigned to blue/red color.
	// Hence we just calculate number for edge that we can add
	blue, red, nocolor := 0, 0, 0
	for _, color := range colors {
		if color == BLUE {
			blue++
		} else if color == RED {
			red++
		} else {
			nocolor++
		}
	}

	// assign uncolored and isolated nodes to
	// blue/red that maximize the additional edges
	if blue > red {
		red += nocolor
	} else {
		blue += nocolor
	}

	return blue*red - len(edges)
}

func main() {
	var numNodes = 5
	var edges = [][2]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 4},
	}
	extraEdge := maxAdditionalEdgeBiparte(edges, numNodes)
	fmt.Println("Additional edge we still can have:", extraEdge)
}
