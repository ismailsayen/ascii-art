package asciiart

import (
	"log"
	"strings"
)

func PrintChar(words []string, slice [][]string) string {
	str := ""
	for _, w := range words { // ici pour mainipuler element par element
		for i := 1; i <= 8; i++ { // pour printer les lignes
			if len(w) == 0 { // ici chaque "" signifie que c'est une newLine
				str += "\n"
				break // ici pour printer \n une seul fois
			}
			for _, e := range w {
				if int(e)-32 >= 0 && int(e)-32 <= len(slice)-1 {
					str += slice[int(e)-32][i] // printer une ligne de chaque lettre du mots
				} else { // si on rencontre un special chartere on retourne une erreur
					log.Fatal("Special charactere is not allowed.")
				}
			}
			if i < 8 { // après chaque ligne retourn a la ligne
				str += "\n"
			}
		}
	}
	if str[len(str)-1] == '\n' && len(words) == 1 {
		return strings.Trim(str, "\n")
	} else if !ContainChar(str) {
		str = str[:len(str)-1]
	}

	return str
}
