package main

//TODO Non gestion des erreurs du fichier

import (
	"fmt"
)

func PrintJose(nbrs_tentative int, Hangman string) {
	var contenuHangmanString []string
	var dessin string
	var compteurRetourLigne int = 0
	//Boucle pour transformer le contenue dans tableau de chaîne de char
	for i := 0; i < len(Hangman); i++ {
		//Récupération de chaque donné du fichier et mettre dans string dessin
		dessin += string(Hangman[i])
		//détection de chaque ligne avec un compteur ligne
		if byte(Hangman[i]) == '\n' {
			compteurRetourLigne++
		}
		// Quand on arrive à 8 lignes on mets notre string dessin dans une nouvelle case
		if compteurRetourLigne == 8 {
			contenuHangmanString = append(contenuHangmanString, dessin)
			dessin = ""
			compteurRetourLigne = 0
		}
	}
	contenuHangmanString = append(contenuHangmanString, dessin)
	//Afficher la case du tableau qui correspond au nbrs de tentative
	fmt.Println(contenuHangmanString[10-nbrs_tentative])
}
