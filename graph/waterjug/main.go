// https://www.geeksforgeeks.org/water-jug-problem-using-bfs/
package main

import "fmt"

func solveWaterJug(jugOneMax, jugTwoMax int, target int) {
	var queue = make([][2]int, 0, 100)
	var paths = make(map[[2]int][2]int)
	var visited = make(map[[2]int]bool)

	queue = append(queue, [2]int{0, 0})

	var pourJugOneTwo = func(jugOne, jugTwo int) [2]int {
		var delta int
		if jugOne < (jugTwoMax - jugTwo) {
			delta = jugOne
		} else {
			delta = jugTwoMax - jugTwo
		}
		return [2]int{jugOne - delta, jugTwo + delta}
	}

	var pourJugTwoOne = func(jugOne, jugTwo int) [2]int {
		var delta int
		if jugTwo < (jugOneMax - jugOne) {
			delta = jugTwo
		} else {
			delta = jugOneMax - jugOne
		}
		return [2]int{jugOne + delta, jugTwo - delta}
	}

	// run breath-first search algorithm
	var goal [2]int
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		// check whether goal state is reached
		// where goal {jugOne, jugTwo} = {0, target}
		jugOne, jugTwo := state[0], state[1]
		if (jugOne == 0 && jugTwo == target) || (jugOne == target && jugTwo == 0) {
			goal = state
			break
		}

		// check if this state was seen before
		if _, seen := visited[state]; seen {
			continue
		}
		visited[state] = true

		// prepare next state that can be traversed from current
		nextStates := [][2]int{
			{jugOneMax, jugTwo},           // fill jug One
			{jugOne, jugTwoMax},           // fill jug two
			{0, jugTwo},                   // empty jug one
			{jugOne, 0},                   // empty jug two
			pourJugOneTwo(jugOne, jugTwo), // pour jug one to two
			pourJugTwoOne(jugOne, jugTwo), // pour jug two to one
		}

		// put all states into queue
		for _, next := range nextStates {
			if _, seen := visited[next]; !seen {
				queue = append(queue, next)
				paths[next] = state
			}
		}
	}

	reversedPaths := make([][2]int, 0, len(paths))
	reversedPaths = append(reversedPaths, goal)
	for state := goal; state != [2]int{0, 0}; {
		state = paths[state]
		reversedPaths = append(reversedPaths, state)
	}

	for i := len(reversedPaths) - 1; i >= 0; i-- {
		fmt.Print(reversedPaths[i], "->")
	}
	fmt.Println("finished")
}

func main() {
	var jugOne, jugTwo = 10, 7
	var target = 9
	solveWaterJug(jugOne, jugTwo, target)
}
