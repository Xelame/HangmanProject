package main

func HideWord(word string, listOfLetterAlreadySay []rune) string {
	hiddenWord := []rune{}        // Initialize hiddenword
	for _, letter := range word { // travel word letter by letter
		isAlreadySay := false                                   // Bool to know the presence of a letter
		for _, letterAppeared := range listOfLetterAlreadySay { // travel letters memories
			if ToUpper(letter) == letterAppeared { // Test if letter does be show
				isAlreadySay = true
			}
		}
		if isAlreadySay { // Show either a letter or dash
			hiddenWord = append(hiddenWord, ToUpper(letter))
			hiddenWord = append(hiddenWord, ' ')
		} else {
			hiddenWord = append(hiddenWord, '_')
			hiddenWord = append(hiddenWord, ' ')
		}
	}
	return string(hiddenWord)
}
