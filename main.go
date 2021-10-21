package main

import (
	"fmt"
	"io/ioutil"
)

// TODO Faire une fonction qui permet de rejouer (PS : réutilisé main en soit ...)
// TODO Add language choose and adapt variables 📚
// FIXME Gestions des constantes
// FIXME Gestion des erreurs
// FIXME Lecture des fichiers (+ entrée joueur)
// FIXME Fractionnage des fichier a r

func main() {
	// Partie recherche du mot dans le fichier
	dictionary, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	wordChoosen := ChooseWord(dictionary)

	// Hidden word part
	lettersAlreadyAppeard := []rune{}
	startHint := wordChoosen[len(wordChoosen)/2-1]
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, ToUpper(rune(startHint)))
	hiddenWord := HideWord(wordChoosen, &lettersAlreadyAppeard)

	//Partie José initialisation
	contenuHangmanByte, err := ioutil.ReadFile("hangman.txt") // TODO Upgrade, if it's possible, the ASCII ART 👨‍🎨
	if err != nil {
		fmt.Println(err.Error())
	}

	//Partie présentation du jeu
	fmt.Println("||Welcome to the Hangman game !             ||\n||Will you be able to find the hidden word ?||")

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
	if attempts != 0 {
		fmt.Println("Well Played you found the word and save Jose !")
	} else {
		Retry()
	}
}

func Retry() {
	fmt.Println("Poor Jose ...\nRetry your chance for him to survive ? [Y]es or [N]o")
	var input string
	fmt.Scanf("%s", &input)
	if len(input) == 1 {
		letter := rune(input[0])
		if ToUpper(letter) == 'Y' {
			main()
		}
	}
}
