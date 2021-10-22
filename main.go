package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// TODO  If it's possible, add language choose and adapt variables 📚
// FIXME Gestions des constantes (notamment les strings) 🕵️‍♂️
// FIXME Lecture des fichiers (+ entrée joueur) 🧙‍♂️
// FIXME Fractionnage des fichier a revoir
// TODO If it's possible, upgrade the ASCII ART 👨‍🎨
//TODO If it's possible, create document for rules and you can open

/*----------------------------------------------------------------------------------------------------------------
**DECLARATION DE CONSTANTE**
 */
const TEXT_ERROR_DICT = "\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||"
const TEXT_ERROR_NO_WORD = "No Word in file"
const TEXT_ERROR_HANG = "\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||"
const TEXT_ERROR_NO_CONTENT = "No content in Hangman.txt"
const TEXT_INTRO = "||Welcome to the Hangman game !                           ||\n||Will you be able to find the hidden word and save José ?||"
const TEXT_FINISH_WIN = "Well Played you found the word and save Jose !\nDo you want to retry ? [Y]es or [N]o"
const TEXT_FINISH_LOST = "Poor Jose ...\nRetry your chance for him to survive ? [Y]es or [N]o"

//----------------------------------------------------------------------------------------------------------------
func main() {
	// Partie recherche du mot dans le fichier
	dictionary, errDict := ioutil.ReadFile("words.txt")
	if errDict != nil {
		errorDectection(errDict.Error())
		log.Fatal(TEXT_ERROR_DICT)
	}
	if string(dictionary) == "" {
		errorDectection(TEXT_ERROR_NO_WORD)
		log.Fatal(TEXT_ERROR_DICT)
	}
	wordChoosen := ChooseWord(dictionary)
	// Hidden word part
	lettersAlreadyAppeard := []rune{}
	startHint := wordChoosen[len(wordChoosen)/2-1]
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint)))
	hiddenWord := HideWord(wordChoosen, &lettersAlreadyAppeard)

	//Partie José initialisation
	contenuHangmanByte, errHang := ioutil.ReadFile("hangman.txt")
	if errHang != nil {
		errorDectection(errHang.Error())
		log.Fatal(TEXT_ERROR_HANG)
	}
	solution := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	if string(contenuHangmanByte) == "" {
		errorDectection(TEXT_ERROR_NO_CONTENT)
		log.Fatal(TEXT_ERROR_HANG)
	}

	//Partie présentation du jeu
	fmt.Println(TEXT_INTRO)

	//Partie boucle de jeu
	attempts := 10
	for isFinished(attempts, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, &lettersAlreadyAppeard))
		PrintJose(attempts, string(contenuHangmanByte)) // Récupération des données du fichier

		// Part Input Player
		GuessingLetter(&lettersAlreadyAppeard)
		if hiddenWord == HideWord(wordChoosen, &lettersAlreadyAppeard) {
			attempts--
		} else {
			hiddenWord = HideWord(wordChoosen, &lettersAlreadyAppeard)
		}
	}
	PrintJose(attempts, string(contenuHangmanByte))
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
	var input string
	fmt.Scanf("%s", &input)
	if len(input) == 1 {
		letter := rune(input[0])
		if ToUpper(letter) == 'Y' {
			main()
		}
	}
}
