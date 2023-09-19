package graph

import (
	"lem-in/models"
)

func NewGraph() *models.Graph {
	return models.NewGraph()
}

func AddRoom(graph *models.Graph, room *models.Room) {
	graph.AddRoom(room)
}

func AddLink(graph *models.Graph, room1, room2 *models.Room) {
	graph.AddLink(room1, room2)
}
