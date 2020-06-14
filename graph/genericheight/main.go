// https://www.geeksforgeeks.org/height-generic-tree-parent-array/
package main

import "fmt"

// computeHeightTreeBFS computes the maximum height of a given tree using breath-first search
// A tree of size n is given as array parent links [0..n-1] where
// every index i in parent[] represents a node and the value at i
// represents the immediate parent of that node.
// For root node value will be -1.
func computeHeightTreeBFS(parentNodeLinks []int) int {
	var queue = make([]int, 0, len(parentNodeLinks))
	var graph = func(parents []int) map[int][]int {
		graphInt := make(map[int][]int)
		for node, parent := range parents {
			if parent == -1 {
				continue
			}
			graphInt[parent] = append(graphInt[parent], node)
		}
		return graphInt
	}(parentNodeLinks)

	queue = append(queue, 0)
	height := -1

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if childs, ok := graph[node]; ok {
				queue = append(queue, childs...)
			}
		}
		height++
	}
	return height
}

// computeHeightTreeDFS computes the maximum height of a given tree
// using depth-first search + memoization/DP.
// A tree of size n is given as array parent links [0..n-1] where
// every index i in parent[] represents a node and the value at i
// represents the immediate parent of that node.
// For root node value will be -1.
func computeHeightTreeDFS(parentNodeLinks []int) int {
	var table = make(map[int]int)
	var dfs func(node int) int

	dfs = func(node int) int {
		if node == -1 {
			return -1
		}

		if height, seen := table[node]; seen {
			return height
		}

		// assume node is marked from [0...n-1]
		height := dfs(parentNodeLinks[node]) + 1
		table[node] = height
		return height
	}

	maxHeight := 0
	for node := 0; node < len(parentNodeLinks); node++ {
		height := dfs(node)
		if height > maxHeight {
			maxHeight = height
		}
	}

	return maxHeight
}

func main() {
	// TEST #1
	parents := []int{-1, 0, 0, 0, 3, 1, 1, 2}
	height := computeHeightTreeBFS(parents)
	fmt.Printf("Test 1: parents = %v\n", parents)
	fmt.Println("Height:", height)

	// TEST #2
	parents = []int{-1, 0, 1, 2, 3}
	height = computeHeightTreeDFS(parents)
	fmt.Printf("Test 2: parents = %v\n", parents)
	fmt.Println("Height:", height)
}
