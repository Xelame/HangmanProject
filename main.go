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
	fmt.Println(hiddenWord)

	//Partie José
	contenuHangmanByte, err := ioutil.ReadFile("hangman.txt") // FIXME Upgrade, if it's possible, the ASCII ART 👨‍🎨
	if err != nil {
		fmt.Println(err.Error())
	}
	PrintJose(8, string(contenuHangmanByte)) // Récupération des données du fichier

	// TODO List of function to show our output (Nathan) 😁
	// TODO Add Guessing function 🤔 (with the HideWord function, he is adapted 👍)
	// TODO Add language choose and adapt variables 📚

	// Part Input Player
	GuessingLetter(&lettersAlreadyAppeard)
	fmt.Print(lettersAlreadyAppeard)
}
