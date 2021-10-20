package main

// Return boolean to know if the game is finished
func isFinished(numberOfAttempts int, word string) bool {
	isRunning := true
	numberOfLetterMissing := 0
	if numberOfAttempts == 0 { // if we haven't no longer attempts
		isRunning = false
	}
	for _, letter := range word { // Count number of underscore
		if letter == '_' {
			numberOfLetterMissing++
		}
	}
	if numberOfLetterMissing == 0 { // if it's remain no more --> all it's appeared
		isRunning = false
	}
	return isRunning
}
