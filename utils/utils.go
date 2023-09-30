package utils

import (
	"fmt"
	"lem-in/choosenantpath"
	"lem-in/movants"
	"lem-in/structure"
	"lem-in/validPaths"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

// Find if string exist in slice
func ExistInSlice(s string, slice []string) bool {
	for _, elem := range slice {
		if elem == s {
			return true
		}
	}
	return false

}

// Check if is number
func IsNumber(n string) bool {
	_, err := strconv.ParseFloat(n, 64)
	return err == nil

}

// Function to count how many time a string is in a slice
func CountElem(s string, tab []string) int {
	c := 0
	for _, elem := range tab {
		if elem == s {
			c++
		}
	}
	return c

}

func FinalPrint(myFarm structure.Farm) string {
	// fmt.Println(myFarm)

	content, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content) + "\n")
	movants.Moveants(choosenantpath.ChoosePath(choosenantpath.CalcPath(validPaths.FindAllPaths(myFarm)), myFarm.AntsNumber), validPaths.FindAllPaths(myFarm))
	c := movants.Moveants(choosenantpath.ChoosePath(choosenantpath.CalcPath(validPaths.FindAllPaths(myFarm)), myFarm.AntsNumber), validPaths.FindAllPaths(myFarm))
	fmt.Println(c)
	return c

}

func RemplacerMotsAvecCoordonnees(texte string, correspondances map[string]structure.Coordinates) string {
	// Créez une expression régulière pour capturer les mots
	regex := regexp.MustCompile(`\b\w+\b`)

	// Utilisez ReplaceAllStringFunc pour effectuer le remplacement en utilisant une fonction de rappel
	nouveauTexte := regex.ReplaceAllStringFunc(texte, func(mot string) string {
		// Recherchez si le mot a une correspondance dans la carte
		if coordonnees, existe := correspondances[mot]; existe {
			// Utilisez les coordonnées pour le remplacement (par exemple, Latitude et Longitude)
			return "ant" + "(" + coordonnees.X + "," + coordonnees.Y + ")"
		}
		// Si le mot n'a pas de correspondance, retournez le mot inchangé
		return mot
	})

	return nouveauTexte
}

func CopieInFile(texte string) {
	err := os.WriteFile("resultvisuliser.txt", []byte(texte), 0644)
	if err != nil {
		fmt.Println("There is an error in the visualiser file.")
	}

}

func RenderTemplate(w http.ResponseWriter, context struct {
	AntsData string
}, templ string) {
	tmpl, err := template.ParseFiles(templ)

	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, context)

	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
		return
	}
}

func GetFirstPart(input string) string {
	parts := strings.Split(input, "-")
	if len(parts) > 0 {
		return parts[0]
	}
	return input // Return the original string if no "-" is found
}
