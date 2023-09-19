package models

import (
	"fmt"
	"strings"
)

type Graph struct {
	Rooms []*Room
	Links map[*Room][]*Room
}

func NewGraph() *Graph {
	return &Graph{
		Rooms: []*Room{},
		Links: make(map[*Room][]*Room),
	}
}

func (g *Graph) AddRoom(room *Room) {
	g.Rooms = append(g.Rooms, room)
}

func (g *Graph) AddLink(roomName1, roomName2 string) error {
	room1 := g.GetRoomByName(roomName1)
	room2 := g.GetRoomByName(roomName2)

	if room1 == nil || room2 == nil {
		return fmt.Errorf("one or both rooms not found")
	}

	g.Links[room1] = append(g.Links[room1], room2)
	g.Links[room2] = append(g.Links[room2], room1)
	return nil
}

func (g *Graph) GetRoomByName(name string) *Room {
	for _, room := range g.Rooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

func (g *Graph) String() string {
	var result strings.Builder

	result.WriteString("Graph Contents:\n")
	for _, room := range g.Rooms {
		result.WriteString(room.String())
		result.WriteString("\nLinked to:")
		for _, linkedRoom := range g.Links[room] {
			result.WriteString(" " + linkedRoom.Name)
		}
		result.WriteString("\n\n")
	}

	return result.String()
}
