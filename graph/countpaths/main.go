// https://www.geeksforgeeks.org/count-possible-paths-two-vertices/
package main

import "fmt"

// countPaths count all possible paths between source and dest vertex
func countPaths(graphInt map[int][]int, numNodes, source, dest int) int {
	var table = make([]int, numNodes)
	var visited = make([]bool, numNodes)
	var dfs func(node int) int

	table[dest] = 1
	visited[dest] = true

	// dfs traverses all possible paths between source and dest vertex
	// using deep first search approach. This method does not work
	// on cyclic graph.
	dfs = func(node int) int {
		if visited[node] {
			return table[node]
		}
		visited[node] = true

		var count = 0
		for _, neig := range graphInt[node] {
			count += dfs(neig)
		}

		table[node] = count
		return count
	}

	pathCount := dfs(source)
	return pathCount
}

func main() {
	var numNodes = 5
	graph := make(map[int][]int)

	// populates directed graph with test edges
	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)
	graph[0] = append(graph[0], 4)
	graph[1] = append(graph[1], 3)
	graph[1] = append(graph[1], 4)
	graph[2] = append(graph[2], 4)
	graph[3] = append(graph[3], 2)

	path04 := countPaths(graph, numNodes, 0, 4)
	path02 := countPaths(graph, numNodes, 0, 2)

	fmt.Printf("There are %d number of paths between node (%d) to (%d)\n", path04, 0, 4)
	fmt.Printf("There are %d number of paths between node (%d) to (%d)\n", path02, 0, 2)
}
