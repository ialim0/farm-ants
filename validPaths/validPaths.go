package validPaths

import (
	"lem-in/structure"
	"sort"
	"strings"
)

// Function to find all matching path for set of ants
func FindAllPaths(graph structure.Farm) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	currentPath := []string{}

	var dfs func(string)
	dfs = func(current string) {
		visited[current] = true
		currentPath = append(currentPath, current)

		if current == graph.EndRoom {
			pathCopy := make([]string, len(currentPath))
			copy(pathCopy, currentPath)
			paths = append(paths, pathCopy)
		} else {
			for _, link := range graph.AllPaths {
				rooms := strings.Split(link, "-")
				if rooms[0] == current && !visited[rooms[1]] {
					dfs(rooms[1])
				} else if rooms[1] == current && !visited[rooms[0]] {
					dfs(rooms[0])
				}
			}
		}

		visited[current] = false
		currentPath = currentPath[:len(currentPath)-1]
	}

	dfs(graph.StartRoom)
	SortPathsByLength(paths)
	return FilterPaths(paths)
}

// Function to sort path
func SortPathsByLength(paths [][]string) {
	sort.SliceStable(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

// Function to filter path to have filtered path
func FilterPaths(paths [][]string) [][]string {
	var filteredPaths [][]string
	usedNodes := make(map[string]bool)

	for _, path := range paths {
		keepPath := true

		for _, node := range path[1 : len(path)-1] {
			if usedNodes[node] {
				keepPath = false
				break
			}
		}

		if keepPath {
			filteredPaths = append(filteredPaths, path)
			// Marquer les nœuds utilisés
			for _, node := range path[1 : len(path)-1] {
				usedNodes[node] = true
			}
		}
	}

	return filteredPaths
}
