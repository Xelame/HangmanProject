/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * game.go                                                             Version : v1.0
 * Programme qui sert à faire la boucle de jeu
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import "fmt"

func Game(attemptsNumber int) {
	//Partie présentation du jeu
	fmt.Println(TEXT_INTRO)
	//Partie boucle de jeu
	for isFinished(attemptsNumber, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, &lettersAlreadyAppeard))
		PrintJose(attemptsNumber, HANGMAN_FILENAME) // Récupération des données du fichier

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
