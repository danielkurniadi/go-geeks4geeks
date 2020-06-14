// https://www.geeksforgeeks.org/clone-an-undirected-graph/
package main

import "fmt"

func cloneGraph(graphInt map[int][]int) map[int][]int {
	var (
		clone   = make(map[int][]int)
		queue   = make([]int, 0, len(graphInt))
		visited = make(map[int]bool)
	)

	for node := range graphInt {
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, seen := visited[node]; seen {
			continue
		}
		visited[node] = true

		var neigs = make([]int, len(graphInt[node]))
		copy(neigs, graphInt[node])
		clone[node] = neigs

		queue = append(queue, graphInt[node]...)
	}

	return clone
}

func prettyPrintGraph(graphInt map[int][]int) {
	fmt.Println("# Graph print:")
	for node := 0; node <= 7; node++ {
		fmt.Printf("(%d): ", node)
		for _, neig := range graphInt[node] {
			fmt.Print(" -> ", neig)
		}
		fmt.Println()
	}
}

func main() {
	graph := make(map[int][]int)

	// populate undirected graph with test edges
	graph[0] = append(graph[0], 1, 4)
	graph[1] = append(graph[1], 0, 2, 3, 4)
	graph[2] = append(graph[2], 1, 3, 6, 7)
	graph[3] = append(graph[3], 1, 2, 4, 5)
	graph[4] = append(graph[4], 0, 1, 3, 5)
	graph[5] = append(graph[5], 2, 3, 4)
	graph[6] = append(graph[6], 2)
	graph[7] = append(graph[7], 2)

	prettyPrintGraph(graph)

	clone := cloneGraph(graph)
	prettyPrintGraph(clone)
}
