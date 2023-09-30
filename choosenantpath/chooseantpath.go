package choosenantpath

import (
	"lem-in/structure"
	"strconv"
)

// Function to help ants choose optimized rooms to move on
func ChoosePath(paths []structure.AntPath, antsnumber int) [][]string {
	var ChoosenRoom = make([][]string, len(paths))

	for i := 1; i <= antsnumber; i++ {
		name := "L" + strconv.Itoa(i)
		for index := range paths {
			if IsMinPath(paths, paths[index].TotalOfAntEroom) {
				ChoosenRoom[index] = append(ChoosenRoom[index], name)
				paths[index].TotalOfAntEroom = paths[index].TotalOfAntEroom + 1
				break
			}
		}
	}
	return ChoosenRoom
}

// Function to have a struct that help choose optimise path for ant
func CalcPath(paths [][]string) []structure.AntPath {
	var tabOfPathData []structure.AntPath
	for index, path := range paths {
		pathData := structure.AntPath{
			PathIndex:       index,
			TotalOfAntEroom: len(path) - 2,
		}
		tabOfPathData = append(tabOfPathData, pathData)
	}
	return tabOfPathData
}

// Funtion to know wich path have minimum step  available
func IsMinPath(Allpaths []structure.AntPath, nbr int) bool {

	for _, path := range Allpaths {
		if nbr > path.TotalOfAntEroom {
			return false
		}

	}
	return true

}
