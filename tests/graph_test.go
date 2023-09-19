// graph/graph_test.go

package tests

import (
	"lem-in/graph"
	"lem-in/models"
	"testing"
)

func TestGraph(t *testing.T) {
	// Create a new graph
	g := graph.NewGraph()

	// Create rooms
	room1 := &models.Room{Name: "Room1", X: 1, Y: 2}
	room2 := &models.Room{Name: "Room2", X: 3, Y: 4}
	room3 := &models.Room{Name: "Room3", X: 5, Y: 6}

	// Add rooms to the graph
	graph.AddRoom(g, room1)
	graph.AddRoom(g, room2)
	graph.AddRoom(g, room3)

	// Add Links between rooms
	graph.AddLink(g, room1, room2)
	graph.AddLink(g, room2, room3)

	// Test if rooms were added
	if len(g.Rooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(g.Rooms))
	}

	// Test if Links were added
	if len(g.Links[room1]) != 1 || len(g.Links[room2]) != 2 || len(g.Links[room3]) != 1 {
		t.Error("Expected Links to be added correctly")
	}
}
