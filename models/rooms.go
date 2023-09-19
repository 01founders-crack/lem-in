package models

type Room struct {
	Name    string
	X, Y    int
	IsStart bool
	IsEnd   bool
}
