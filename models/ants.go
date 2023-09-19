package models

// Ant represents an ant in the ant farm.
type Ant struct {
	ID          int   // Unique identifier for the ant
	CurrentRoom *Room // Current room the ant is in
}

// NewAnt creates a new Ant instance with the given ID and starting room.
func NewAnt(id int, startRoom *Room) *Ant {
	return &Ant{
		ID:          id,
		CurrentRoom: startRoom,
	}
}

// MoveTo moves the ant to a new room.
func (ant *Ant) MoveTo(newRoom *Room) {
	ant.CurrentRoom = newRoom
}

// IsEmpty checks if the room is empty (no ants in the room).
func (room *Room) IsEmpty() bool {
	// Implement your logic to check if the room is empty here
	return true // Modify this based on your implementation
}

// PlaceAnt places an ant in the room.
func (room *Room) PlaceAnt(ant *Ant) {
	// Implement logic to add an ant to the room here
}

// ReleaseAnt removes an ant from the room.
func (room *Room) ReleaseAnt() {
	// Implement logic to remove an ant from the room here
}
