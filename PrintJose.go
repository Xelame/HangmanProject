package main

//TODO Non gestion des erreurs du fichier

import (
	"fmt"
	"io/ioutil"
)

func PrintJosé(nbrs_tentative int, Hangman string) {
	var contenuHangmanString []string
	var dessin string
	var compteurRetourLigne int = 0
	var contenuHangmanByte []byte
	contenuHangmanByte, _ = ioutil.ReadFile(Hangman) // Récupération des données du fichier
	//Boucle pour transformer le contenue dans tableau de chaîne de char
	for i := 0; i < len(contenuHangmanByte); i++ {
		//Récupération de chaque donné du fichier et mettre dans string dessin
		dessin += string(contenuHangmanByte[i])
		//détection de chaque ligne avec un compteur ligne
		if rune(contenuHangmanByte[i]) == '\n' {
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
