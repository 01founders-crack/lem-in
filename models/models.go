package models

type Graph struct {
	Rooms []*Room
}

type Room struct {
	Name    string
	X, Y    int
	IsStart bool
	IsEnd   bool
}
