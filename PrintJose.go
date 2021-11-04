/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * main.go
 * Programme qui sert à afficher José l'homme pendu                     Version : v1.0
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

const HANGMAN_LINE = 8

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func PrintJose(nbrs_tentative int, hangmanFileName string) {
	//Partie José initialisation
	contenuHangmanByte, errOpen := os.Open(hangmanFileName)
	if errOpen != nil {
		errorDectection(errOpen.Error())
		log.Fatal(TEXT_ERROR_HANG)
	}
	defer contenuHangmanByte.Close()

	reader := bufio.NewScanner(contenuHangmanByte)

	for i := 1; reader.Scan(); i++ {
		if (10-nbrs_tentative)*HANGMAN_LINE <= i && i <= (10-nbrs_tentative+1)*HANGMAN_LINE {
			fmt.Println(reader.Text())
		}
	}
}
