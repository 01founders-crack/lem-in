package tests

import (
	"lem-in/helpers"
	"lem-in/models"
	"testing"
)

func TestParseRoom(t *testing.T) {
	roomStr := "Room1 1 2"
	expectedRoom := &models.Room{
		Name:    "Room1",
		X:       1,
		Y:       2,
		IsStart: false,
		IsEnd:   false,
	}

	room := helpers.ParseRoom(roomStr)

	if room == nil {
		t.Error("Expected a room, got nil")
		return
	}

	if room.Name != expectedRoom.Name ||
		room.X != expectedRoom.X ||
		room.Y != expectedRoom.Y ||
		room.IsStart != expectedRoom.IsStart ||
		room.IsEnd != expectedRoom.IsEnd {
		t.Errorf("Expected room %+v, got %+v", expectedRoom, room)
	}
}

func TestParseRooms(t *testing.T) {
	roomStrs := []string{
		"Room1 1 2",
		"Room2 3 4",
		"Room3 5 6",
	}

	parsedRooms := helpers.ParseRooms(roomStrs)

	if len(parsedRooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(parsedRooms))
	}
}

// Add more test cases as needed
