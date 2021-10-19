package main

import (
	"fmt"
	"io/ioutil"
)

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
	lettersAlreadyAppeard = append(lettersAlreadyAppeard, rune(startHint))
	hiddenWord := HideWord(wordChoosen, lettersAlreadyAppeard)

	//Partie José initialisation
	contenuHangmanByte, err := ioutil.ReadFile("hangman.txt") // FIXME Upgrade, if it's possible, the ASCII ART 👨‍🎨
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO Add language choose and adapt variables 📚

	//Partie présentation du jeu
	fmt.Println("||Welcome to the Hangman game !             ||\n||Will you be able to find the hidden word ?||")
	//Partie boucle principale
	attempts := 0
	for isFinish(8, hiddenWord) {
		fmt.Println(HideWord(wordChoosen, lettersAlreadyAppeard))
		PrintJose(attempts, string(contenuHangmanByte)) // Récupération des données du fichier

		// Part Input Player
		GuessingLetter(&lettersAlreadyAppeard)
		if hiddenWord == HideWord(wordChoosen, lettersAlreadyAppeard) {
			attempts--
		} else {
			hiddenWord = HideWord(wordChoosen, lettersAlreadyAppeard)
		}
	}
	if attempts == 0 {
		fmt.Println("Poor José .... Retry your chance for him to survive ?")
	}
}
