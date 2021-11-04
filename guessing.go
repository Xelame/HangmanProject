/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 19/10/21
 * guessing.go                                                          Version : v1.0
 * Programme qui sert à gerer la partie des choix du joueur (lettre)
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import "fmt"

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

const ALREADY_SAYS = "This letter is already says."
const ASK_LETTER = "Can you give me a letter please : "
const WRITE_SOMETHING = "Don't write something else than a letter please."
const ONE_LETTER = "Don't write more than one letter please."

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var letterGuessed rune = ' '
var isValid bool = true
var count int = 0
var guessingLetter rune = ' '

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

// TODO Add Comments 😉 !!!

func GuessingLetter(listOfLetterAlreadySay *[]rune) {
	fmt.Print(ASK_LETTER)
	for fmt.Scanf("%s", &input); !IsNotValid(input, *listOfLetterAlreadySay); fmt.Scanf("%s", &input) {
		fmt.Print(ASK_LETTER)
	}
	for _, value := range input {
		letterGuessed = ToUpper(rune(value))
	}
	*listOfLetterAlreadySay = append(*listOfLetterAlreadySay, letterGuessed)
}

func IsNotValid(guessingInput string, listOfLetterAlreadySay []rune) bool {
	count = 0
	isValid = true
	for range guessingInput {
		count++
	}
	if count == 1 {
		guessingLetter = ToUpper(rune(guessingInput[0]))
		for _, letterAlreadyHere := range listOfLetterAlreadySay {
			if guessingLetter == letterAlreadyHere {
				isValid = false
				fmt.Println(ALREADY_SAYS)
			}
		}
		if !(IsUpper(guessingLetter) || IsExctendedAsciiLetter(guessingLetter)) {
			isValid = false
			fmt.Println(WRITE_SOMETHING)
		}
	} else {
		isValid = false
		fmt.Println(ONE_LETTER)
	}
	return isValid
}

func IsUpper(value rune) bool {
	if 'A' <= int(value) && int(value) <= 'Z' {
		return true
	}
	return false
}

func ToUpper(value rune) rune {
	if ('a' <= value && value <= 'z') || (224 <= value && value <= 255) {
		value -= 32
	}
	return value
}

func IsExctendedAsciiLetter(value rune) bool {
	if 192 <= value && value <= 255 {
		return true
	}
	return false
}
