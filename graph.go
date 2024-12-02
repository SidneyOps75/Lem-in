package main

import (
	"container/list"
	"errors"
)

func FindShortestPaths(farm *AntFarm) ([][]string, error) {
	// Find start and end rooms
	var startRoom, endRoom *Room
	for name, room := range farm.Rooms {
		if room.IsStart {
			startRoom = room
			farm.Start = name
		}
		if room.IsEnd {
			endRoom = room
			farm.End = name
		}
	}

	if startRoom == nil || endRoom == nil {
		return nil, errors.New("missing start or end room")
	}

	// BFS to find shortest paths
	paths := [][]string{}
	type PathInfo struct {
		Room    string
		Path    []string
		Visited map[string]bool
	}

	queue := list.New()
	queue.PushBack(PathInfo{
		Room:    farm.Start,
		Path:    []string{farm.Start},
		Visited: map[string]bool{farm.Start: true},
	})

	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(PathInfo)

		if current.Room == farm.End {
			paths = append(paths, current.Path)
			continue
		}

		currentRoom := farm.Rooms[current.Room]
		for neighbor := range currentRoom.Neighbors {
			if !current.Visited[neighbor] {
				newPath := make([]string, len(current.Path))
				copy(newPath, current.Path)
				newPath = append(newPath, neighbor)

				newVisited := make(map[string]bool)
				for k, v := range current.Visited {
					newVisited[k] = v
				}
				newVisited[neighbor] = true

				queue.PushBack(PathInfo{
					Room:    neighbor,
					Path:    newPath,
					Visited: newVisited,
				})
			}
		}
	}

	if len(paths) == 0 {
		return nil, errors.New("no path exists between start and end")
	}

	return paths, nil
}
