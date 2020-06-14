// https://www.geeksforgeeks.org/iterative-depth-first-traversal/
package main

import "fmt"

func iterativeDFS(graphInt map[int][]int, numNodes int, source int) {
	var stack = make([]int, 0, numNodes)
	var visited = make([]bool, numNodes)

	stack = append(stack, source)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] == true {
			continue
		}

		visited[node] = true
		fmt.Print(node, "-> ")

		for _, neigh := range graphInt[node] {
			stack = append(stack, neigh)
		}

	}
	fmt.Print("\n")
}

func main() {
	var graph = make(map[int][]int)

	// populates
	graph[0] = append(graph[0], 1)
	graph[1] = append(graph[1], 2)
	graph[2] = append(graph[2], 3)
	graph[3] = append(graph[3], 4)
	graph[4] = append(graph[4], 0)
	iterativeDFS(graph, 5, 0)
}
