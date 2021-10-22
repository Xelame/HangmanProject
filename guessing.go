package main

import "fmt"

// TODO Add Comments 😉 !!!

func GuessingLetter(listOfLetterAlreadySay *[]rune) {
	var letterGuessed rune
	fmt.Print("Can you give me a letter please : ")
	for fmt.Scanf("%s", &input); !IsNotValid(input, *listOfLetterAlreadySay); fmt.Scanf("%s", &input) {
		fmt.Print("Can you give me a letter please : ")
	}
	letterGuessed = ToUpper(rune(input[0]))
	*listOfLetterAlreadySay = append(*listOfLetterAlreadySay, letterGuessed)
}

func IsNotValid(guessingInput string, listOfLetterAlreadySay []rune) bool {
	isValid := true
	if len(guessingInput) == 1 {
		guessingLetter := ToUpper(rune(guessingInput[0]))
		for _, letterAlreadyHere := range listOfLetterAlreadySay {
			if guessingLetter == letterAlreadyHere {
				isValid = false
				fmt.Println("This letter is already says.")
			}
		}
		if !(IsUpper(guessingLetter) || IsLower(guessingLetter)) {
			isValid = false
			fmt.Println("Don't write something else than a letter please.")
		}
	} else {
		isValid = false
		fmt.Println("Don't write more than one letter please.")
	}
	return isValid
}

func IsLower(value rune) bool {
	if 'a' <= int(value) && int(value) <= 'z' {
		return true
	}
	return false
}

func IsUpper(value rune) bool {
	if 'A' <= int(value) && int(value) <= 'Z' {
		return true
	}
	return false
}

func ToUpper(value rune) rune {
	if 'a' <= value && value <= 'z' {
		value -= 32
	}
	return value
}
