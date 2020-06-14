// https://www.geeksforgeeks.org/shortest-path-reach-one-prime-changing-single-digit-time/
package main

import (
	"fmt"
	"math"
)

func shortestDigitChangeBFS(graphPrime map[int][]int, source int, dest int) {
	if source == dest {
		return
	}

	var queue = make([]int, 0, 100)
	var paths = make(map[int]int)
	var visited = make([]bool, 9999)

	queue = append(queue, source)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == dest {
			break
		}

		if visited[node] == true {
			continue
		}

		visited[node] = true
		for _, neig := range graphPrime[node] {
			if visited[neig] == true {
				continue
			}
			queue = append(queue, neig)
			paths[neig] = node
		}
	}

	// backtrack paths to shortest prime digit change
	fmt.Println(paths)
}

func makePrimesGraph(primes []int) map[int][]int {
	graph := make(map[int][]int)

	for _, prime := range primes {
		graph[prime] = make([]int, 0, 10)
	}

	for _, prime := range primes {
		for i, x := 0, prime; x > 0; i, x = i+1, x/10 {
			mult := int(math.Pow(float64(10), float64(i)))
			digit := mult * (x % 10)
			
			for j := 0; j < 10; j++ {
				num := prime - digit + j*mult
				if _, ok := graph[num]; ok {
					graph[prime] = append(graph[prime], num)
					graph[num] = append(graph[num], prime)
				}
			}
		}
	}
	return graph
}

func sieveOfEratosthenes(n int) []int {
	var primes = make([]int, 0, n)
	var notPrime = make([]bool, n+1)

	for p := 2; p*p <= n; p++ {
		if notPrime[p] == false { // prime number
			for i := p * p; i < n+1; i += p {
				notPrime[i] = true
			}
		}
	}
	for p := 1000; p < n+1; p++ {
		if notPrime[p] == false { // prime number
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	primes := sieveOfEratosthenes(9999)
	graph := makePrimesGraph(primes)

	// run and test shortest prime digit discovery
	// > 1033, 8179
	// > 1373, 8017
	// > 1033, 1033
	source, dest := 1033, 8179
	shortestDigitChangeBFS(graph, source, dest)
}
