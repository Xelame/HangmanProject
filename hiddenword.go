package main

import (
	"fmt"
)

// TODO Validité des entrées & test tout en minuscule et affichage tout en majuscule

func HideWord(word string, listOfLetterAlreadySay *[]rune) string {
	hiddenWord := []rune{}        // Initialize hiddenword
	for _, letter := range word { // travel word letter by letter
		isAlreadySay := false                                    // Bool to know the presence of a letter
		for _, letterAppeared := range *listOfLetterAlreadySay { // travel letters memories
			if letter == letterAppeared { // Test if letter does be show
				isAlreadySay = true
			}
		}
		if isAlreadySay { // Show either a letter or dash
			hiddenWord = append(hiddenWord, letter)
			hiddenWord = append(hiddenWord, ' ')
		} else {
			hiddenWord = append(hiddenWord, '_')
			hiddenWord = append(hiddenWord, ' ')
		}
	}
	return string(hiddenWord)
}

// TODO Add Comments 😉

func GuessingLetter(listOfLetterAlreadySay *[]rune) {
	var input string
	var letterGuessed rune
	fmt.Print("Can you give me a letter please : ")
	for fmt.Scanf("%s", &input); !IsNotValid(input, *listOfLetterAlreadySay); fmt.Scanf("%s", &input) {
		fmt.Print("Can you give me a letter please : ")
	}
	letterGuessed = rune(input[0])
	*listOfLetterAlreadySay = append(*listOfLetterAlreadySay, rune(letterGuessed))
}

func IsNotValid(guessingInput string, listOfLetterAlreadySay []rune) bool {
	isValid := true
	if len(guessingInput) == 1 {
		guessingLetter := rune(guessingInput[0])
		for _, letterAlreadyHere := range listOfLetterAlreadySay {
			if guessingLetter == letterAlreadyHere {
				isValid = false
				fmt.Println("This letter is already says.")
			}
		}
		if !(('A' <= guessingLetter && guessingLetter <= 'Z') || ('a' <= guessingLetter && guessingLetter <= 'z')) {
			isValid = false
			fmt.Println("Don't write something else than a letter please.")
		}
	} else {
		isValid = false
		fmt.Println("Don't write more than one letter please.")
	}
	return isValid
}
