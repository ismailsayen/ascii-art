package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	asciiart "asciiart/Functions"
)

func main() {
	// Vérifie si l'utilisateur a passé exactement un argument.
	if len(os.Args) != 2 {
		fmt.Println("please enter a word")
		return
	}

	// Ouvre le fichier contenant les définitions des symboles.
	file, err := os.Open("./Files/standard.txt")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer file.Close() // Assure la fermeture du fichier après usage

	sentence := os.Args[1]

	// Vérifie que le mot n'est pas vide
	if len(sentence) == 0 {
		log.Fatal("Erreur: Please enter a Word.")
	}

	scanner := bufio.NewScanner(file) // Initialise un scanner pour lire le fichier ligne par ligne

	count := 0
	symbole := []string{}
	symboles := [][]string{}

	// Boucle à travers chaque ligne du fichier texte
	for scanner.Scan() {
		line := scanner.Text()
		symbole = append(symbole, line)
		count++

		// Chaque symbole ASCII dans ce format est constitué de 9 lignes
		if count == 9 {
			symboles = append(symboles, symbole) // Ajoute le symbole complet à la liste de symboles
			symbole = []string{}
			count = 0
		}
	}
	// Vérifie si le fichier contient au moins 95 symboles.
	if len(symboles) < 95 {
		log.Fatal("Please make sure all characters are present in the file.")
	}

	result := asciiart.PrintChar(asciiart.Split(sentence), symboles)

	// Affiche le résultat
	fmt.Println(result)
}
