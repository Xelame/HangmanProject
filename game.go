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
	"os"
	"os/exec"
)

const HANGMAN_FILENAME = "hangman.txt"
const DICTIONARY_FILENAME = "words.txt"
const TEXT_ERROR_DICT = "\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||"
const TEXT_ERROR_NO_WORD = "No Word in file"
const TEXT_ERROR_HANG = "\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||"
const TEXT_ERROR_NO_CONTENT = "No content in hangman file"
const TEXT_FINISH_WIN = "Well Played you found the word and save Jose !\nDo you want to retry ? \033[92m[Y]es or \033[31m[N]o\033[0m"
const TEXT_FINISH_LOST = "Poor Jose ...\nRetry your chance for him to survive ? \033[92m[Y]es\033[0m or \033[31m[N]o\033[0m"

var input string = ""
var solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}

var lettersAlreadyAppeard = []rune{'-'}

func Game(attemptsNumber int) {
	var wordChoosen string = ChooseWord(DICTIONARY_FILENAME)
	var startHint byte = wordChoosen[len(wordChoosen)/2-1]
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint)))
	var hiddenWord string = HideWord(wordChoosen, lettersAlreadyAppeard)
	for isFinished(attemptsNumber, hiddenWord) {
		Clear()
		fmt.Println(HideWord(wordChoosen, lettersAlreadyAppeard))
		PrintJose(attemptsNumber, HANGMAN_FILENAME) // Récupération des données du fichier
		AttemptsColor(attemptsNumber)
		// Part Input Player
		GuessingLetter()
		if hiddenWord == HideWord(wordChoosen, lettersAlreadyAppeard) {
			attemptsNumber--
		} else {
			hiddenWord = HideWord(wordChoosen, lettersAlreadyAppeard)
		}

	}
	PrintJose(attemptsNumber, HANGMAN_FILENAME)
	fmt.Println(HideWord(wordChoosen, solution))
	if attemptsNumber != 0 {
		Animation(winText)
		fmt.Println(TEXT_FINISH_WIN)
	} else {
		fmt.Println(TEXT_FINISH_LOST)
	}
}

func Retry() {
	fmt.Scanf("%s", &input)
	Clear()
	if len(input) == 1 {
		letter := rune(input[0])
		if ToUpper(letter) == 'Y' {
			lettersAlreadyAppeard = []rune{'-'}
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

func Clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
