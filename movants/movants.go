package movants

import (
	"fmt"
	"strings"
)

// Function to move all ants throught all rooms
func Moveants(antsslice, roomsslice [][]string) string {
	var movetab []string
	for index, ant := range antsslice {
		c := strings.Split(MoveOneAnts(ant, roomsslice[index]), "\n")
		movetab = MergeSlices(movetab, c)

	}

	return strings.Join(movetab, "\n")

}

// Funtion to move ants throught giving rooms
func MoveOneAnts(ants []string, rooms []string) string {
	c := ""
	state := make(map[string]string)
	state[ants[0]] = rooms[0]
	for _, ant := range ants[1:] {
		state[ant] = ""
	}

	for step := 1; step < len(ants)+len(rooms)-1; step++ {

		for _, ant := range ants {
			currentRoom := state[ant]
			if currentRoom == "" {
				continue
			}

			nextRoom := getNextRoom(currentRoom, rooms)
			state[ant] = nextRoom
			if !strings.Contains(c, fmt.Sprintf(" %s-%s", ant, nextRoom)) {

				c += fmt.Sprintf(" %s-%s", ant, nextRoom)

			}

		}

		for _, ant := range ants {
			if state[ant] == "" {
				state[ant] = rooms[0]
				break
			}
		}
		if step != len(ants)+len(rooms)-2 {
			c += "\n"

		}

	}
	return c

}

// Fonction to move to next room
func getNextRoom(currentRoom string, rooms []string) string {
	for i, room := range rooms {
		if room == currentRoom {
			if i+1 < len(rooms) {
				return rooms[i+1]
			}
		}
	}
	return currentRoom
}

// Function to merge two slice
func MergeSlices(slice1 []string, slice2 []string) []string {
	var mergedSlice []string

	minLength := len(slice1)
	if len(slice2) < minLength {
		minLength = len(slice2)
	}

	for i := 0; i < minLength; i++ {
		mergedSlice = append(mergedSlice, slice1[i]+slice2[i])
	}

	if len(slice1) > len(slice2) {
		mergedSlice = append(mergedSlice, slice1[minLength:]...)
	} else {
		mergedSlice = append(mergedSlice, slice2[minLength:]...)
	}

	return mergedSlice
}
