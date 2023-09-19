package helpers

import (
	"lem-in/models"
	"strconv"
	"strings"
)

func ParseRoom(str string) *models.Room {
	roomInfo := strings.Fields(str)
	if len(roomInfo) == 3 {
		x, _ := strconv.Atoi(roomInfo[1])
		y, _ := strconv.Atoi(roomInfo[2])
		return &models.Room{
			Name:    roomInfo[0],
			X:       x,
			Y:       y,
			IsStart: false,
			IsEnd:   false,
		}
	}
	return nil
}

func ParseRooms(rooms []string) []*models.Room {
	var parsedRooms []*models.Room

	for _, roomStr := range rooms {
		room := ParseRoom(roomStr)
		if room != nil {
			parsedRooms = append(parsedRooms, room)
		}
	}

	return parsedRooms
}
