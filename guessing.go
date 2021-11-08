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
// Partie du programme
// -----------------------------------------------------------------------------------

// TODO Add Comments 😉 !!!

func GuessingLetter() {
	var letterGuessed rune = ' '
	fmt.Print(ASK_LETTER)
	for fmt.Scanf("%s", &input); IsValidEntry(input); fmt.Scanf("%s", &input) {
		fmt.Print(ASK_LETTER)
	}
	for _, value := range input {
		letterGuessed = ToUpper(rune(value))
	}
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, letterGuessed)
}

func IsValidEntry(guessingInput string) bool {
	var guessingLetter rune = ' '
	var count int = 0
	isNotValid := false
	for range guessingInput {
		count++
	}
	if count == 1 {
		guessingLetter = ToUpper(rune(guessingInput[0]))
		for _, letterAlreadyHere := range lettersAlreadyAppeard {
			if guessingLetter == letterAlreadyHere {
				isNotValid = true
				fmt.Println(ALREADY_SAYS)
			}
		}
		if !(IsUpper(guessingLetter) || IsExctendedAsciiLetter(guessingLetter)) {
			isNotValid = true
			fmt.Println(WRITE_SOMETHING)
		}
	} else {
		isNotValid = true
		fmt.Println(ONE_LETTER)
	}
	return isNotValid
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
