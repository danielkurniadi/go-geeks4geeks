// https://www.geeksforgeeks.org/print-paths-given-source-destination-using-bfs/
package main

import "fmt"

// computePaths computes and stores all possible paths
// between source and destination vertex. Graph may contain cycle.
// Graph can be directed and undirected graph.
func computePaths(graphInt map[int][]int, source, dest int) [][]int {
	var (
		queue = make([][]int, 0, len(graphInt))
		paths = make([][]int, 0, len(graphInt))
	)

	queue = append(queue, []int{source})

	for len(queue) > 0 {
		// enqueue state (path) from queue
		path := queue[0]
		queue = queue[1:]
		if len(path) == 0 {
			continue
		}
		// the last node in the path is current the node
		node := path[len(path)-1]

		// check if we reach destination
		if node == dest {
			paths = append(paths, path)
			continue
		}

		// traverse the next nodes
		for _, neig := range graphInt[node] {
			// TODO: use hashset(map) to check cycle
			if sliceHasElement(path, neig) {
				continue
			}
			newPath := append(path, neig)
			queue = append(queue, newPath)
		}
	}
	return paths
}

func sliceHasElement(slice []int, element int) bool {
	for i := len(slice) - 1; i >= 0; i-- {
		value := slice[i]
		if value == element {
			return true
		}
	}
	return false
}

func prettyPrintPaths(paths [][]int) {
	for _, path := range paths {
		fmt.Printf("(%d) ", path[0])
		for _, node := range path[1:] {
			fmt.Print(" -> ", node)
		}
		fmt.Println()
	}
}

func main() {
	var graph = make(map[int][]int)

	graph[0] = append(graph[0], 1, 2, 3)
	graph[1] = append(graph[1], 3)
	graph[2] = append(graph[2], 0, 1)

	source, dest := 2, 3
	paths := computePaths(graph, source, dest)
	fmt.Println(paths)
}
