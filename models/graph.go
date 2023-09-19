package models

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

func (g *Graph) AddLink(room1, room2 *Room) {
	g.Links[room1] = append(g.Links[room1], room2)
	g.Links[room2] = append(g.Links[room2], room1)
}
