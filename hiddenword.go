/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * guessing.go                                                          Version : v1.0
 * Programme qui sert à gerer la partie des choix du joueur (lettre)
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var hiddenWord = []rune{} // Initialize hiddenword
var isAlreadySay bool = false

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func HideWord(word string, listOfLetterAlreadySay *[]rune) string {
	for _, letter := range word { // travel word letter by letter                                   // Bool to know the presence of a letter
		for _, letterAppeared := range *listOfLetterAlreadySay { // travel letters memories
			if ToUpper(letter) == letterAppeared { // Test if letter does be show
				isAlreadySay = true
			}
		}
		if isAlreadySay { // Show either a letter or dash
			hiddenWord = append(hiddenWord, ToUpper(letter))
			hiddenWord = append(hiddenWord, ' ')
		} else {
			hiddenWord = append(hiddenWord, '_')
			hiddenWord = append(hiddenWord, ' ')
		}
	}
	return string(hiddenWord)
}
