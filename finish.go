/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 19/10/21
 * finish.go
 * Programme qui détecte si le jeu Hangman est fini                     Version : v1.0
 * ---------------------------------------------------------------------------------*/
package main

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var isRunning bool = true
var numberOfLetterMissing int = 0

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func isFinished(numberOfAttempts int, word string) bool {
	isRunning = true
	numberOfLetterMissing = 0
	if numberOfAttempts == 0 { // if we haven't no longer attempts
		isRunning = false
	}
	for _, letter := range word { // Count number of underscore
		if letter == '_' {
			numberOfLetterMissing++
		}
	}
	if numberOfLetterMissing == 0 { // if it's remain no more --> all it's appeared
		isRunning = false
	}
	return isRunning
}
