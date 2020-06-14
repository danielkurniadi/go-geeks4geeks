// https://www.geeksforgeeks.org/find-a-mother-vertex-in-a-graph/
package main

import "fmt"

func deepFirstSearch(node int, visited []bool, graphInt map[int][]int) {
	if visited[node] == true {
		return
	}

	visited[node] = true
	deepFirstSearch(node, visited, graphInt)

	if neighs, ok := graphInt[node]; ok {
		for _, neigh := range neighs {
			deepFirstSearch(neigh, visited, graphInt)
		}
	}

	return
}

func findMotherVertex(numNodes int, graphInt map[int][]int) (mother int) {
	var visited = make([]bool, numNodes)

	fmt.Println(visited)
	// traverse graph, assume nodes are 0 to N-1
	for node := 0; node < numNodes; node++ {
		if visited[node] == true {
			continue
		}
		deepFirstSearch(node, visited, graphInt)
		mother = node
	}

	// check if mother is truly mother vertex
	// see reference:
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
	deepFirstSearch(mother, visited, graphInt)

	for i, visit := range visited {
		if visit == false {
			fmt.Println(i)
			return -1
		} // no mother vertex
	}

	return
}

func main() {
	var numNodes = 7

	graph := make(map[int][]int)
	for node := range graph {
		graph[node] = make([]int, 0, numNodes)
	}

	// populate directed graph with test nodes
	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)
	graph[1] = append(graph[1], 3)
	graph[4] = append(graph[4], 1)
	graph[6] = append(graph[6], 4)
	graph[5] = append(graph[5], 2)
	graph[5] = append(graph[5], 6)
	graph[6] = append(graph[6], 0)

	mother := findMotherVertex(numNodes, graph)
	fmt.Println("Mother Vertex:", mother)
}
