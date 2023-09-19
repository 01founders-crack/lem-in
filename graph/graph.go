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

func AddLink(graph *models.Graph, roomName1, roomName2 string) error {
	return graph.AddLink(roomName1, roomName2)
}
