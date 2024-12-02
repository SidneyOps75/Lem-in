package main

import (
	"fmt"
	"sort"
	"strings"
)

func SimulateAntMovement(farm *AntFarm, paths [][]string) []string {
	// Sort paths by length for optimal distribution
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	turns := []string{}
	antLocations := make(map[int]string)
	usedPaths := make(map[string]bool)
	antCount := farm.Ants

	currentTurn := 0
	for antCount > 0 {
		currentMoves := []string{}
		currentPathIndex := 0

		for ant := 1; ant <= farm.Ants; ant++ {
			// Skip ants already in end room
			if antLocations[ant] == farm.End {
				continue
			}

			// Place or move ant
			if antLocations[ant] == "" {
				// First turn: place ant on first path
				currentPath := paths[currentPathIndex]
				antLocations[ant] = currentPath[1]
				currentMoves = append(currentMoves, fmt.Sprintf("L%d-%s", ant, antLocations[ant]))
				usedPaths[currentPath[0]+"-"+currentPath[1]] = true
				currentPathIndex = (currentPathIndex + 1) % len(paths)
			} else {
				// Find current room in current path
				for _, path := range paths {
					for i := 1; i < len(path); i++ {
						if antLocations[ant] == path[i-1] {
							nextRoom := path[i]
							// Ensure next room isn't occupied
							roomOccupied := false
							for _, loc := range antLocations {
								if loc == nextRoom {
									roomOccupied = true
									break
								}
							}

							if !roomOccupied {
								antLocations[ant] = nextRoom
								currentMoves = append(currentMoves, fmt.Sprintf("L%d-%s", ant, nextRoom))
								break
							}
						}
					}
				}
			}
		}

		// Check if all ants reached end
		allReachedEnd := true
		for ant := 1; ant <= farm.Ants; ant++ {
			if antLocations[ant] != farm.End {
				allReachedEnd = false
				break
			}
		}

		if len(currentMoves) > 0 {
			turns = append(turns, sortMoves(currentMoves))
		}

		if allReachedEnd {
			break
		}

		currentTurn++
		if currentTurn > farm.Ants*len(paths[0]) {
			break
		}
	}

	return turns
}

// sortMoves ensures moves are sorted lexicographically
func sortMoves(moves []string) string {
	sort.Strings(moves)
	return strings.Join(moves, " ")
}
