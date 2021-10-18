package main

// TODO Add Comment ^^

// TODO Ready to be use for another function 😉
func HideWord(word string, listOfLetterAlreadySay []rune) string {
	hiddenWord := []rune{}
	for _, letter := range word {
		isAlreadySay := false
		for _, letterAppeared := range listOfLetterAlreadySay {
			if letter == letterAppeared {
				isAlreadySay = true
			}
		}
		if isAlreadySay {
			hiddenWord = append(hiddenWord, letter)
			hiddenWord = append(hiddenWord, ' ')
		} else {
			hiddenWord = append(hiddenWord, '_')
			hiddenWord = append(hiddenWord, ' ')
		}
	}
	return string(hiddenWord)
}
