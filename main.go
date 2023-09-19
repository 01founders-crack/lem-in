package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file for reading. Replace "input.txt" with your actual file name.
	file, err := os.Open("examples/example01.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file.
	for scanner.Scan() {
		line := scanner.Text()

		// Check for lines starting with "##start" and "##end" and print them.
		if strings.HasPrefix(line, "##start") {
			fmt.Println("Start:", line)
		} else if strings.HasPrefix(line, "##end") {
			fmt.Println("End:", line)
		} else {
			// For other lines, split them by '-' and print the parts.
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				fmt.Printf("From: %s, To: %s\n", parts[0], parts[1])
			} else {
				fmt.Println("Invalid line format:", line)
			}
		}
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}
