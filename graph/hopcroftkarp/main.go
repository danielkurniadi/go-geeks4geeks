package main

import "math"

// HopcroftKarpBiparteMatching implements Hopcroft Karp algorithm
// for finding maximum matching for unweighted undirected bipartiate graphs
// https://en.wikipedia.org/wiki/Hopcroft%E2%80%93Karp_algorithm
// Finds maximum matching for unweighted bipartiate graph in O(E sqrt(v)).
/// The Hopcroft Karp algorithm is as follows:
/// 1.  Augmented path has end points as two unmatched vertices.
///     There may be 0 or more other vertices in the path besides source and destination however they must all be matched vertices.
///     The edges in the path must alternate between unmatched edge and matched edge. As the first vertice in path is unmatched, first edge is
///     automatically unmatched. If next vertex is unmatched as well then path ends otherwise we are on matched vertex in which case we must follow
///     edge from previous matching. After we follow the matched edge, we would be on other side but now we can not have any other matched edge so
///     automatically, next edge must be unmatched to go on opposite side. And so the path continues.
/// 2. Shortest augmented path is just that: Shortest in length of all augmented paths in graph.
/// 3. There may be multiple shortest augmented paths of the same length.
/// 4. Hopcroft Karp algorithm requires that we construct the maximal set of shortest augmented paths that don't have any common vertex between them.
/// 5. Then so symmetric difference of existing matching and all augmented paths to get the new matching.
/// 6. Repeat until no more augmented paths are found.

// constants
const (
	NIL = -1
	INF = math.MaxInt32
)

// HopcroftKarpBiparteMatching implements Hopcroft Karp algorithm
// for finding maximum matching for unweighted undirected bipartiate graphs.
// It inputs a graph representation (map int to slide of int) and
// returns a map of node U to V that yield maximum matched nodes.
func HopcroftKarpBiparteMatching(bipGraphInt map[int][]int, countU int, countV int) {

}

// hopcroftkarpBFSHelper traverses the biparte graph in breatd-first search
// and updates the distances between the dummy node (NIL) to node in vertexes U.
// It then returns true when shortest path to dummy node (NIL) is found
func hopcroftkarpBFSHelper(dists, pairU, pairV map[int]int, bipGraphInt map[int][]int) bool {
	var queue = make([]int, 0, len(pairU))

	for u := range bipGraphInt {
		if pairU[u] == NIL {
			dists[u] = 0
			// Enqueue all unmatched vertices in queue
			queue = append(queue, u)
			continue
		}
		dists[u] = INF
	}

	// Set distance of dummy node to infinite. When we find the shortest path, we would end at dummy node and BFS
	// would set its distance to the length of shortest path. We can use this length to eliminate any path that are
	// longer than. If more than one vertex has same length shortest path then for both of them we can follow the
	// dist array all the way to dummy node and when we get there we can check that length of path so far is the
	// same value as in dist[NIL].
	dists[NIL] = INF

	for len(queue) > 0 {
		// Dequeue next node
		u := queue[0]
		queue = queue[1:]

		// If length of path to this node exceeds the shortest path
		// means we already found this path and now we ignore this path
		if dists[u] < dists[NIL] {
			for _, v := range bipGraphInt[u] {
				nextU := pairV[v]

				if dists[nextU] == INF {
					dists[nextU] = dists[u] + 1

					// Note that queue will always contains vertice from set of U.
					// We don't need dists array for V because from vertex in U we always
					// go to next vertex in U
					queue = append(queue, nextU)
				}
			}
		}
	}
	// If we found a shortest path, then distance to dummy node (dists[NIL]) would
	// contain length of the shortest path
	return dists[NIL] != INF
}

// hopcroftkarpDFSHelper traverses the unmatched node in depth-first search manner
// using the value in dists. If we are able to arrive at dummy node (NIL) using vertex u
// and the distance from vertex u and vertex v is one, we found one of the shortest paths
func hopcroftkarpDFSHelper(u int, dists, pairU, pairV map[int]int, bipGraphInt map[int][]int) bool {
	if u == NIL {
		return true
	}
	for _, v := range bipGraphInt[u] {
		nextU := pairV[v]

		// Our neighbor path is the next node to traverse
		// if it's matching node is our distance +1
		if (dists[nextU] == dists[u]+1) || hopcroftkarpDFSHelper(nextU, dists, pairU, pairV, bipGraphInt) {
			pairU[u] = v
			pairV[v] = u
			return true
		}
	}
	return false
}
