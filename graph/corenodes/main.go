// https://www.geeksforgeeks.org/find-k-cores-graph/
package main

import "fmt"

// computeGraphKCores computes the k-cores version of undirected graph
// by pruning nodes that has edge degree less than k.
// Computation is inplace and modifying original graph
func computeGraphKCores(k int, graphInt map[int][]int, numNodes int) map[int][]int {
	var queue = make([]int, 0, 100)

	for node := 0; node < numNodes; node++ {
		if neighs, ok := graphInt[node]; ok {
			if len(neighs) < k {
				queue = append(queue, node)
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neigh := range graphInt[node] {
			graphInt[neigh] = deleteElement(graphInt[neigh], node)
			if len(graphInt[neigh]) < k {
				queue = append(queue, neigh)
			}
		}

		delete(graphInt, node)
	}
	return graphInt
}

// deleteElement deletes a value from slice. If there are two or more
// elements equal to target value, only the first one will be removed
func deleteElement(slice []int, target int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		if slice[i] == target {
			slice[i], slice[n-1] = slice[n-1], slice[i]
			slice = slice[:n-1]
			break
		}
	}

	return slice
}

func prettyPrintGraph(numNodes int, graphInt map[int][]int) {
	fmt.Println("K Cores Graph:")

	for node := 0; node < numNodes; node++ {
		fmt.Printf("(%d) -> ", node)
		if neighs, ok := graphInt[node]; ok {
			for _, neigh := range neighs {
				fmt.Print(neigh, "-> ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var numNodes = 9
	graph := make(map[int][]int)

	// populate undirected graph with test nodes
	graph[0] = append(graph[0], 1)
	graph[0] = append(graph[0], 2)

	graph[1] = append(graph[1], 0)
	graph[1] = append(graph[1], 2)
	graph[1] = append(graph[1], 5)

	graph[2] = append(graph[2], 0)
	graph[2] = append(graph[2], 1)
	graph[2] = append(graph[2], 3)
	graph[2] = append(graph[2], 4)
	graph[2] = append(graph[2], 5)
	graph[2] = append(graph[2], 6)

	graph[3] = append(graph[3], 2)
	graph[3] = append(graph[3], 4)
	graph[3] = append(graph[3], 6)
	graph[3] = append(graph[3], 7)

	graph[4] = append(graph[4], 2)
	graph[4] = append(graph[4], 3)
	graph[4] = append(graph[4], 6)
	graph[4] = append(graph[4], 7)

	graph[5] = append(graph[5], 1)
	graph[5] = append(graph[5], 2)
	graph[5] = append(graph[5], 6)
	graph[5] = append(graph[5], 8)

	graph[6] = append(graph[6], 2)
	graph[6] = append(graph[6], 3)
	graph[6] = append(graph[6], 5)
	graph[6] = append(graph[6], 7)
	graph[6] = append(graph[6], 8)

	graph[7] = append(graph[7], 3)
	graph[7] = append(graph[7], 4)
	graph[7] = append(graph[7], 6)

	graph[8] = append(graph[8], 5)
	graph[8] = append(graph[8], 6)

	// compute kcores graph
	graph = computeGraphKCores(3, graph, numNodes)

	// print result
	prettyPrintGraph(numNodes, graph)
}
