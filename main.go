package main

import (
	"fmt"
	"lem-in/helpers"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		
	// Open the input file for reading. Replace "input.txt" with your actual file name.
	file, err := os.Open("examples/"+os.Args[1])
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	helpers.ScanFileReturnAnts(file)
	}
	
}
