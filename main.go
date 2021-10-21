package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// FIXME Gestion des erreurs
// FIXME Changez le traveling des string
// TODO  If it's possible, add language choose and adapt variables 📚
// FIXME Gestions des constantes (notamment les strings)
// FIXME Fractionnage des fichier a revoir
// TODO If it's possible, upgrade the ASCII ART 👨‍🎨

func main() {
	// Partie recherche du mot dans le fichier
	dictionary, errDict := ioutil.ReadFile("words.txt")
	if errDict != nil {
		errorDectection(errDict.Error())
		log.Fatal("\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||")
	}
	if string(dictionary) == "" {
		errDictNoWord := "No Word in file"
		errorDectection(errDictNoWord)
		log.Fatal("\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||")
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
		log.Fatal("\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||")

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	solution := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	if string(contenuHangmanByte) == "" {
		errDictNoWord := "No content in Hangman.txt"
		errorDectection(errDictNoWord)
		log.Fatal("\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||")
	}

	//Partie présentation du jeu
	fmt.Println("||Welcome to the Hangman game !             ||\n||Will you be able to find the hidden word ?||")

	//Partie boucle de jeu
	attempts := 10
	for isFinished(attempts, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, &lettersAlreadyAppeard))
		PrintJose(attempts, scanner.Text()) // Récupération des données du fichier

		// Part Input Player
		GuessingLetter(&lettersAlreadyAppeard)
		if hiddenWord == HideWord(wordChoosen, &lettersAlreadyAppeard) {
			attempts--
		} else {
			hiddenWord = HideWord(wordChoosen, &lettersAlreadyAppeard)
		}
	}
	PrintJose(attempts, scanner.Text())
	fmt.Println(HideWord(wordChoosen, &solution))
	if attempts != 0 {
		fmt.Println("Well Played you found the word and save Jose !\nDo you want to retry ? [Y]es or [N]o")
		Retry()
	} else {
		fmt.Println("Poor Jose ...\nRetry your chance for him to survive ? [Y]es or [N]o")
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
