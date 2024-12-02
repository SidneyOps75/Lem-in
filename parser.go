package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name      string
	X, Y      int
	IsStart   bool
	IsEnd     bool
	Neighbors map[string]bool
}

type AntFarm struct {
	Ants          int
	Rooms         map[string]*Room
	Start, End    string
	OriginalInput string
	RawRooms      []string
	RawLinks      []string
}

func ParseInput(filename string) (*AntFarm, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	farm := &AntFarm{
		Rooms:    make(map[string]*Room),
		RawRooms: []string{},
		RawLinks: []string{},
	}

	scanner := bufio.NewScanner(file)
	var fileContent strings.Builder
	var parsingPhase int // 0: ants, 1: rooms, 2: links
	var nextIsStart, nextIsEnd bool

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fileContent.WriteString(line + "\n")

		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		switch parsingPhase {
		case 0: // Parsing number of ants
			//fmt.Printf("Parsing number of ants: %s\n", line)
			ants, err := strconv.Atoi(line)
			if err != nil {
				return nil, errors.New("invalid number of ants")
			}
			farm.Ants = ants
			parsingPhase = 1

		case 1: // Parsing rooms
			//fmt.Printf("Parsing room line: %s\n", line)

			// Handle start/end commands
			if line == "##start" {
				nextIsStart = true
				continue
			}
			if line == "##end" {
				nextIsEnd = true
				continue
			}

			// Handle room definitions
			if strings.Contains(line, "-") {
				parsingPhase = 2
				farm.RawLinks = append(farm.RawLinks, line)
				continue
			}

			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				x, err1 := strconv.Atoi(parts[1])
				y, err2 := strconv.Atoi(parts[2])
				if err1 != nil || err2 != nil {
					return nil, errors.New("invalid room coordinates")
				}

				if parts[0][0] == 'L' || parts[0][0] == '#' {
					return nil, errors.New("invalid room name")
				}

				room := &Room{
					Name:      parts[0],
					Neighbors: make(map[string]bool),
					IsStart:   nextIsStart,
					IsEnd:     nextIsEnd,
					X:         x,
					Y:         y,
				}

				farm.Rooms[parts[0]] = room
				farm.RawRooms = append(farm.RawRooms, line)

				if nextIsStart {
					//fmt.Printf("Setting start room: %s\n", parts[0])
					farm.Start = parts[0]
					nextIsStart = false
				}
				if nextIsEnd {
					//fmt.Printf("Setting end room: %s\n", parts[0])
					farm.End = parts[0]
					nextIsEnd = false
				}
			}

		case 2: // Parsing links
			//fmt.Printf("Parsing link: %s\n", line)
			farm.RawLinks = append(farm.RawLinks, line)
		}
	}

	// Process links
	//fmt.Printf("Processing links\n")
	for _, link := range farm.RawLinks {
		rooms := strings.Split(link, "-")
		if len(rooms) != 2 {
			return nil, errors.New("invalid link format")
		}

		room1, ok1 := farm.Rooms[rooms[0]]
		room2, ok2 := farm.Rooms[rooms[1]]

		if !ok1 || !ok2 {
			return nil, errors.New("link references unknown room")
		}

		room1.Neighbors[rooms[1]] = true
		room2.Neighbors[rooms[0]] = true
	}

	farm.OriginalInput = fileContent.String()

	// Validate start and end rooms are set
	if farm.Start == "" || farm.End == "" {
		fmt.Printf("Start room: %s\nEnd room: %s\n", farm.Start, farm.End)
		return nil, errors.New("missing start or end room")
	}

	return farm, nil
}
