// https://www.geeksforgeeks.org/count-number-nodes-given-level-using-bfs/
package main

import "fmt"

func levelOrderBFS(graphInt map[int][]int, numNodes int, source int) []int {
	var queue = make([]int, 0, numNodes)
	var visited = make([]bool, numNodes)

	if _, ok := graphInt[source]; !ok {
		return []int{}
	}

	queue = append(queue, source)
	levels := make([]int, 0, 100)
	for len(queue) > 0 {
		n := len(queue)

		// save number of nodes in this level
		levels = append(levels, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if visited[node] == true {
				continue
			}
			if neighs, ok := graphInt[node]; ok {
				queue = append(queue, neighs...)
			}
		}
	}
	return levels
}

func main() {
	var numNodes = 7
	var graph = make(map[int][]int, 0)

	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)
	graph[1] = append(graph[1], 3)
	graph[1] = append(graph[1], 4)
	graph[1] = append(graph[1], 5)
	graph[2] = append(graph[2], 6)

	levels := levelOrderBFS(graph, numNodes, 0)

	// print count of nodes for all levels
	for level, count := range levels {
		fmt.Printf("Levels (%d) has %d nodes", level, count)
	}
}
