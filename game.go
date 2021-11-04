/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * game.go                                                             Version : v1.0
 * Programme qui sert à faire la boucle de jeu
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import (
	"fmt"
)

func Game(attemptsNumber int) {
	fmt.Print(HANGMAN_BANNER)
	//Partie présentation du jeu
	fmt.Println(TEXT_INTRO)
	//Partie boucle de jeu
	for isFinished(attemptsNumber, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, &lettersAlreadyAppeard))
		PrintJose(attemptsNumber, HANGMAN_FILENAME) // Récupération des données du fichier
		AttemptsColor(attemptsNumber)
		// Part Input Player
		GuessingLetter(&lettersAlreadyAppeard)
		if hiddenWord == HideWord(wordChoosen, &lettersAlreadyAppeard) {
			attemptsNumber--
		} else {
			hiddenWord = HideWord(wordChoosen, &lettersAlreadyAppeard)
		}

	}
	PrintJose(attemptsNumber, HANGMAN_FILENAME)
	fmt.Println(HideWord(wordChoosen, &solution))
	if attemptsNumber != 0 {
		fmt.Println(TEXT_FINISH_WIN)
	} else {
		fmt.Println(TEXT_FINISH_LOST)
	}
}

func Retry() {
	fmt.Scanf("%s", &input)
	if len(input) == 1 {
		letter := rune(input[0])
		if ToUpper(letter) == 'Y' {
			wordChoosen = ChooseWord(DICTIONARY_FILENAME)
			startHint = wordChoosen[len(wordChoosen)/2-1]
			lettersAlreadyAppeard = []rune{'-'}
			lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint)))
			main()
		}
	}
}

func AttemptsColor(attemptsNumber int) {
	switch {
	case 8 <= attemptsNumber && attemptsNumber <= 10:
		fmt.Printf("You have \033[32m%d\033[0m tries left\n", attemptsNumber)
	case 5 <= attemptsNumber && attemptsNumber <= 7:
		fmt.Printf("You have \033[33m%d\033[0m tries left\n", attemptsNumber)
	case 2 <= attemptsNumber && attemptsNumber <= 4:
		fmt.Printf("You have \033[31m%d\033[0m tries left\n", attemptsNumber)
	case attemptsNumber == 1:
		fmt.Printf("You have \033[31m%d\033[0m try left\n", attemptsNumber)
	}

}
