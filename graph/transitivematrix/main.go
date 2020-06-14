// https://www.geeksforgeeks.org/transitive-closure-of-a-graph-using-dfs/
package main

import "fmt"

func prettyPrintMatrix(matrix [][]bool) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	return
}

func transitiveMatrix(numNodes int, graphInt map[int][]int) (matrix [][]bool) {
	matrix = make([][]bool, numNodes)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]bool, numNodes)
	}

	var dfs func(int, int)

	dfs = func(source, target int) {
		if matrix[source][target] == true {
			return
		}

		matrix[source][target] = true
		for _, neig := range graphInt[target] {
			dfs(source, neig)
		}
	}

	for node := 0; node < numNodes; node++ {
		dfs(node, node)
	}

	return matrix
}

func main() {
	var numNodes = 4
	var graph = make(map[int][]int)

	for node := 0; node < numNodes; node++ {
		graph[node] = make([]int, 0, numNodes)
	}

	// populate directed graph with test edges
	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)
	graph[1] = append(graph[1], 2)
	graph[2] = append(graph[2], 0)
	graph[2] = append(graph[2], 3)
	graph[3] = append(graph[3], 3)

	// print results
	transmat := transitiveMatrix(numNodes, graph)
	prettyPrintMatrix(transmat)
	return
}
