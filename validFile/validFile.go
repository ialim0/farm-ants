package validFile

import (
	"fmt"
	"lem-in/structure"
	"lem-in/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

var allpaths, roomName []string
var ValidRoom [][]string
var DataStr string
var Coord map[string]structure.Coordinates

// Function to valid file
func ValidFile(fileName string) {
	coordinatesMap := make(map[string]structure.Coordinates)
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)

	}
	strContent := string(content)
	tabContent := strings.Split(strContent, "\n")
	nbrInt, _ := strconv.Atoi(tabContent[0])
	if !utils.IsNumber(tabContent[0]) || nbrInt < 1 {
		printError("ERROR: invalid data format, invalid number of Ants")

	} else if utils.CountElem("##start", tabContent) != 1 || utils.CountElem("##end", tabContent) != 1 {
		printError("ERROR: invalid data format, ##start or ##end format")

	} else {
		var roomInfos []string

		for _, elem := range tabContent {

			if tabContent[0] != elem && !strings.HasPrefix(elem, "#") && len(elem) != 0 {

				roomtab := strings.Split(elem, " ")
				path := strings.Split(elem, "-")

				if len(roomtab) == 3 || len(path) == 2 {
					if len(roomtab) == 3 {
						if utils.ExistInSlice(elem, roomInfos) {
							printError("ERROR: invalid data format, room can't start with L")

						}
						if strings.Contains(elem, " ") {
							temp := roomtab[1] + " " + roomtab[2]
							if utils.ExistInSlice(temp, roomInfos) {
								printError("ERROR: invalid data format,room can't have same coordonnees.")

							} else {
								roomInfos = append(roomInfos, temp)

							}

						}

						if strings.HasPrefix(roomtab[0], "L") {
							fmt.Println("ERROR: invalid data format, room can't start with L")
							os.Exit(1)

						}
						if !IsValidCoord(roomtab[1:]) {
							printError("ERROR: invalid data format, the room identifiers error")
						}
						if utils.ExistInSlice(roomtab[0], roomName) {
							printError("ERROR: invalid data format, Can't have room with same name.")

						} else {
							roomName = append(roomName, roomtab[0])
							coordinatesMap[roomtab[0]] = structure.Coordinates{X: roomtab[1], Y: roomtab[2]}
							ValidRoom = append(ValidRoom, roomtab)
						}

					} else {

						if !IsValidPath(path, roomName) {
							printError("ERROR: invalid data format, link to unkown room or room link to himself or link line error")
						}
						if strings.Contains(elem, "-") {
							allpaths = append(allpaths, elem)

						}

					}

				} else {
					printError("ERROR: invalid data format path")

				}

			}

		}

	}
	myFarm := structure.Farm{
		AntsNumber: nbrInt,
		StartRoom:  StartEndRoom("##start", tabContent),
		EndRoom:    StartEndRoom("##end", tabContent),
		AllPaths:   allpaths,
		AllRooms:   roomName,
	}
	DataStr = utils.FinalPrint(myFarm)
	Coord = coordinatesMap

	DataStr = utils.RemplacerMotsAvecCoordonnees(DataStr, coordinatesMap)
	utils.CopieInFile(DataStr)
	// Print final result utils.FinalPrint(myFarm)

}

// Function to have starter room from graph
func StartEndRoom(roomType string, tabcontent []string) string {
	for index, elem := range tabcontent {
		if index != 0 && tabcontent[index-1] == roomType {
			return strings.Split(elem, " ")[0]
		}
	}
	return ""
}

// Function to valid coordonnees
func IsValidCoord(tab []string) bool {
	for _, elem := range tab {
		if !utils.IsNumber(elem) {
			return false
		}

	}
	return true

}

// Function to check status of path in the file
func IsValidPath(tabpath, tabRoomName []string) bool {
	for _, elem := range tabpath {
		if !utils.ExistInSlice(elem, tabRoomName) || utils.CountElem(elem, tabpath) != 1 {
			return false

		}

	}
	return true

}

// Function to print error of file validation step
func printError(errormessage string) {
	fmt.Println(errormessage)
	os.Exit(1)

}
