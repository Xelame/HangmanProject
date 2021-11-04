// TODO  If it's possible, add language choose and adapt variables 📚
// TODO If it's possible, create document for rules and you can open
// FIXME Problème dans retry ça ne change pas de mots
// FIXME Gestions des constantes (notamment les strings) 🕵️‍♂️
// FIXME Lecture des fichiers (+ entrée joueur) 🧙‍♂️
// FIXME Fractionnage des fichier a revoir
// FIXME Retravailler le desing pour plsu explicite

/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * main.go
 * Programme principale                                                 Version : v1.0
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import (
	"fmt"
)

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

const HANGMAN_FILENAME = "hangman.txt"
const DICTIONARY_FILENAME = "words.txt"
const TEXT_ERROR_DICT = "\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||"
const TEXT_ERROR_NO_WORD = "No Word in file"
const TEXT_ERROR_HANG = "\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||"
const TEXT_ERROR_NO_CONTENT = "No content in hangman file"
const TEXT_INTRO = "||Welcome to the Hangman game !                           ||\n||Will you be able to find the hidden word and save José ?||"
const TEXT_FINISH_WIN = "Well Played you found the word and save Jose !\nDo you want to retry ? [Y]es or [N]o"
const TEXT_FINISH_LOST = "Poor Jose ...\nRetry your chance for him to survive ? [Y]es or [N]o"

var solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var wordChoosen string = ""
var input string = ""

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func main() {
	wordChoosen = ChooseWord(DICTIONARY_FILENAME)
	// Hidden word part
	lettersAlreadyAppeard := []rune{'-'}           //VAR
	startHint := wordChoosen[len(wordChoosen)/2-1] //VAR
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint)))
	hiddenWord := HideWord(wordChoosen, &lettersAlreadyAppeard) //VAR

	//Partie présentation du jeu
	fmt.Println(TEXT_INTRO)

	//Partie boucle de jeu
	attempts := 10 //VAR
	for isFinished(attempts, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, &lettersAlreadyAppeard))
		PrintJose(attempts, HANGMAN_FILENAME) // Récupération des données du fichier

		// Part Input Player
		GuessingLetter(&lettersAlreadyAppeard)
		if hiddenWord == HideWord(wordChoosen, &lettersAlreadyAppeard) {
			attempts--
		} else {
			hiddenWord = HideWord(wordChoosen, &lettersAlreadyAppeard)
		}
	}
	PrintJose(attempts, HANGMAN_FILENAME)
	fmt.Println(HideWord(wordChoosen, &solution))
	if attempts != 0 {
		fmt.Println(TEXT_FINISH_WIN)
		Retry()
	} else {
		fmt.Println(TEXT_FINISH_LOST)
		Retry()
	}
}

func Retry() {
	fmt.Scanf("%s", &input)
	if len(input) == 1 {
		letter := rune(input[0]) //VAR
		if ToUpper(letter) == 'Y' {
			main()
		}
	}
}
