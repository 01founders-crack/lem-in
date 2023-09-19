package main

import (
	"fmt"
	"lem-in/graph"
	"lem-in/helpers"
	"os"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {

		// Open the input file for reading. Replace "input.txt" with your actual file name.
		file, err := os.Open("examples/" + os.Args[1])
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		_, Rooms, Links := helpers.ScanFileReturnAnts(file)
		graph1 := graph.NewGraph()
		AllRooms := helpers.ParseRooms(Rooms)
		for i := 0; i < len(AllRooms); i++ {
			graph.AddRoom(graph1, AllRooms[i])
		}
		for i := 0; i < len(Links); i++ {
			tempLink := strings.Split(Links[i], "-")
			graph.AddLink(graph1, tempLink[0], tempLink[1])
		}
		fmt.Println(graph1)
	}
}
