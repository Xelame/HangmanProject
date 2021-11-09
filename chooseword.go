/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * chooseword.go                                                        Version : v1.0
 * Programme qui sert choisir un mot aléatoire parmi le dictionnaire
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import (
	"math/rand"
	"time"
)

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var randomNumber int = 0
var contentOfDictionary = []string{}

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func ChooseWord(dictionary string) string {

	scanner := OpenScanner(dictionary)

	for scanner.Scan() {
		contentOfDictionary = append(contentOfDictionary, scanner.Text())
	}
	randomNumber = ChooseRandomNumber(len(contentOfDictionary)) // Generate random number
	return ReadWord(contentOfDictionary, randomNumber)          // Return my word
}

// Function read the word in my dictionary
func ReadWord(fileContent []string, wordPositionChoosen int) string {
	//Partie José initialisation
	var myWord string

	for index, fileWord := range fileContent {
		if index == wordPositionChoosen {
			myWord = fileWord
		}
	}
	return myWord // Return the word
}

// Function to choose a random number beetween 0 and the last word position
func ChooseRandomNumber(numberOfWords int) int {
	// Init Source type with a seed who is always different because the number is affiliated with the clock
	randomSource := rand.NewSource(time.Now().UnixNano())
	// Init Rand type who have a random number in each type
	randomValue := rand.New(randomSource)
	// Fix a limit for the random number who is my number of word
	return randomValue.Intn(numberOfWords)
}
