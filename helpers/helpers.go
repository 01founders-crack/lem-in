package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ContainsStartCharCheck(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstChar := s[0] // First byte, not necessarily the first rune (character)
	if firstChar == '#' {
		return true
	} else if firstChar == 'L' {
		return true
	}
	return false
}

func ContainsSpace(s string) bool {
	for _, char := range s {
		if char == ' ' {
			return true
		}
	}
	return false
}

func ScanFileReturnAnts(file *os.File)(int, [][]string, [][]string) {
	// Create a scanner to read lines from the file.
	scanner := bufio.NewScanner(file)
	var links [][]string
	var rooms [][]string
	var temprooms [][]string
	var startRoom []string
	var endRoom []string

	TotalAnts := 0
	mindex := 0         // Initialize line index
	IsEnd := false      // Initialize isEnd check
	IsStart := false    // init start check
	isFirstLine := true // Initialize isFirstLine check
	// Loop through each line in the file.
	for scanner.Scan() {
		line := scanner.Text()
		if isFirstLine {
			ants, _ := strconv.Atoi(line)
			TotalAnts = ants
			fmt.Println("Total ants:", TotalAnts)
			isFirstLine = false
		} else if strings.HasPrefix(line, "##start") {
			fmt.Println("Start:", line)
			IsStart = true
			// Check for lines starting with "##start" and "##end" and print them.
		} else if strings.HasPrefix(line, "##end") {
			fmt.Println("End:", line)
			IsEnd = true
		} else if strings.Contains(line, "-") {
			// For other lines, split them by '-' and print the parts.
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				links = append(links, parts)
				fmt.Printf("From: %s, To: %s\n", parts[0], parts[1])
			} else {
				fmt.Println("Invalid line format:", line)
			}
		} else {
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				if IsStart {
					IsStart = false
					// add to the Start
					startRoom = parts
					fmt.Printf("START FOUND Name: %s, X: %s, Y: %s\n", parts[0], parts[1], parts[2])
				} else if IsEnd {
					IsEnd = false
					// add to the end
					endRoom = parts
					fmt.Printf("END FOUND Name: %s, X: %s, Y: %s\n", parts[0], parts[1], parts[2])
				} else {
					// add to array
					temprooms = append(temprooms, parts)
					fmt.Printf("Name: %s, X: %s, Y: %s\n", parts[0], parts[1], parts[2])
				}

			} else {
				fmt.Println("Invalid line format:", line)
			}
		}
		mindex++
	}

	// Check for any errors during scanning.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	// Concatenate slices
	rooms = append(rooms, startRoom)    // Adding startRoom to rooms
	rooms = append(rooms, temprooms...) // Adding temprooms to rooms
	rooms = append(rooms, endRoom)      // Adding endRoom to rooms

	// Print concatenated slice
	for _, room := range links {
		fmt.Println(room)
	}
	return TotalAnts, rooms, links
}
