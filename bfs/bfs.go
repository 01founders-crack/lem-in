package bfs

import "fmt"

// Pseudo-code for BFS
func BFS(graph Graph, start, end Room) []Room {
	visited := make(map[Room]bool)
	queue := []Room{start}
	parent := make(map[Room]Room)

	for len(queue) > 0 {
		currentRoom := queue[0]
		queue = queue[1:]

		if currentRoom == end {
			break // Found the shortest path
		}

		for _, neighbor := range graph.Links[currentRoom] {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = currentRoom
				queue = append(queue, neighbor)
			}
		}
	}

	// Reconstruct the shortest path
	shortestPath := []Room{end}
	for current := end; current != start; current = parent[current] {
		shortestPath = append([]Room{parent[current]}, shortestPath...)
	}

	return shortestPath
}

// Pseudo-code for moving ants along the path
func MoveAnts(ants []Ant, shortestPaths [][]Room) {
	for turn := 0; turn < len(shortestPaths[0]); turn++ {
		for antIndex, ant := range ants {
			currentRoom := ant.CurrentRoom
			nextRoom := shortestPaths[antIndex][turn]

			// Check if the next room is empty
			if nextRoom.IsEmpty() {
				ant.MoveTo(nextRoom)
				currentRoom.ReleaseAnt()
				nextRoom.PlaceAnt(ant)
				// Print the ant's movement
				fmt.Printf("L%d-%s ", ant.ID, nextRoom.Name)
			}
		}
		fmt.Println() // Print a newline for the next turn
	}
}

// Parse input, build graph, and initialize ants
// Run BFS to find shortest paths
// Move ants along the paths and print their movements
