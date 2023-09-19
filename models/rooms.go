package models

import "fmt"

type Room struct {
	Name    string
	X, Y    int
	IsStart bool
	IsEnd   bool
}

func (r *Room) String() string {
	startEnd := ""
	if r.IsStart {
		startEnd = " (Start)"
	} else if r.IsEnd {
		startEnd = " (End)"
	}
	return fmt.Sprintf("Room: %s (%d, %d)%s", r.Name, r.X, r.Y, startEnd)
}
