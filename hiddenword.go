package main

import (
	"bufio"
	"fmt"
	"os"
)

//FIXME case en trop !

// TODO Add Comment ^^

// TODO Ready to be use for another function 😉
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

func GuessingLetter(listOfLetterAlreadySay *[]rune) {
	var text string
	for len(text) != 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Give me a letter : ")
		text, _ = reader.ReadString('\n')
	}
	*listOfLetterAlreadySay = append(*listOfLetterAlreadySay, rune(text[0]))
}
