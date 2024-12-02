package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if file is provided
	if len(os.Args) < 2 {
		fmt.Println("ERROR: No input file provided")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Parse the input file
	farm, err := ParseInput(filename)
	if err != nil {
		fmt.Printf("ERROR: invalid data format - %v\n", err)
		os.Exit(1)
	}

	// Find optimal paths
	paths, err := FindShortestPaths(farm)
	if err != nil {
		fmt.Printf("ERROR: path finding failed - %v\n", err)
		os.Exit(1)
	}

	// Simulate ant movement
	moves := SimulateAntMovement(farm, paths)

	// Print input file content
	fmt.Println(farm.OriginalInput)

	// Print ant moves
	for _, turn := range moves {
		fmt.Println(turn)
	}
}
