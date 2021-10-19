package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO Validité des entrées & test tout en minuscule et affichage tout en majuscule

func HideWord(word string, listOfLetterAlreadySay []rune) string {
	hiddenWord := []rune{}        // Initialize hiddenword
	for _, letter := range word { // travel word letter by letter
		isAlreadySay := false                                   // Bool to know the presence of a letter
		for _, letterAppeared := range listOfLetterAlreadySay { // travel letters memories
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
	var letterGuessed rune
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Can you give me a letter please : ")
	for input, err := reader.ReadString('\n'); !IsNotValid(input, *listOfLetterAlreadySay); input, err = reader.ReadString('\n') {
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print("Can you give me a letter please : ")
		letterGuessed = rune(input[0])
	}
	*listOfLetterAlreadySay = append(*listOfLetterAlreadySay, rune(letterGuessed))
}

func IsNotValid(guessingInput string, listOfLetterAlreadySay []rune) bool {
	isValid := true
	if len(guessingInput) == 2 {
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
