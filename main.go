package main

import (
	"fmt"
	"lem-in/source/parse"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Program takes argument (fileName)")
		os.Exit(1)
	}
	argument := os.Args[1]
	parse.RunProgramWithFile(argument, false)
}
