package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"lem-in/validFile"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ant struct {
	Name      string `json:"name"`
	Movements []struct {
		X    int `json:"x"`
		Y    int `json:"y"`
		Step int `json:"step"`
	} `json:"movements"`
}
type ConsolidatedAnt struct {
	Name      string `json:"name"`
	Movements []struct {
		X    int `json:"x"`
		Y    int `json:"y"`
		Step int `json:"step"`
	} `json:"movements"`
}

//

func main() {

	if len(os.Args) == 2 {
		fileName := os.Args[1]
		validFile.ValidFile(fileName)
		port := ":3030"
		// Open the file for reading
		file, err := os.Open("resultvisuliser.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		// Initialize antsData slice
		var antsData []Ant
		// Regular expression to match the ant data format
		antPattern := regexp.MustCompile(`L\d+-\w+\((\d+),(\d+)\)`)
		var currentAnt Ant
		scanner := bufio.NewScanner(file)
		currentStep := 0
		for scanner.Scan() {
			line := scanner.Text()
			// Split the line by spaces to get individual ant movements
			antMovements := strings.Fields(line)
			for _, antMovement := range antMovements {
				// Match the ant movement against the regular expression
				match := antPattern.FindStringSubmatch(antMovement)
				if match != nil {
					// Extract ant name, X, and Y coordinates
					antName := strings.Split(match[0], "(")[0]
					x, _ := strconv.Atoi(match[1])
					y, _ := strconv.Atoi(match[2])
					// Check if the ant name matches the current ant or if it's a new ant
					if antName != currentAnt.Name {
						// Add the previous ant to the slice if it's not the first ant
						if currentAnt.Name != "" {
							antsData = append(antsData, currentAnt)
						}
						// Initialize a new ant
						currentAnt = Ant{Name: antName}
						currentAnt.Movements = []struct {
							X    int `json:"x"`
							Y    int `json:"y"`
							Step int `json:"step"`
						}{}
					}
					// Add the movement to the current ant with the current step
					currentAnt.Movements = append(currentAnt.Movements, struct {
						X    int `json:"x"`
						Y    int `json:"y"`
						Step int `json:"step"`
					}{X: x, Y: y, Step: currentStep})
				}
			}
			// Increment the step for the next ant
			currentStep++
		}
		// Append the last ant to the slice
		if currentAnt.Name != "" {
			antsData = append(antsData, currentAnt)
		}
		// consolid := consolidateAnts(antsData)
		// Print the antsData in JSON format
		// antsJSON, _ := json.MarshalIndent(consolid, "", "  ")
		// // fmt.Println(string(antsJSON))
		// Serve the HTML file and ants data via HTTP
		http.Handle("/", http.FileServer(http.Dir("./static"))) // Serve static files from the current directory
		http.HandleFunc("/ants", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			// Encode and write antsData as JSON to the response
			if err := json.NewEncoder(w).Encode(antsData); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
		// Start the HTTP server
		fmt.Printf("Server listening on port %s\n", port)
		if err := http.ListenAndServe(port, nil); err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
	}

}
