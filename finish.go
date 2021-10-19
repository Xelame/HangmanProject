package main

// TODO Add Comments ^^
func isFinish(numberOfAttempts int, word string) bool {
	isRunning := true
	numberOfLetterMissing := 0
	if numberOfAttempts == 0 {
		isRunning = false
	}
	for _, letter := range word {
		if letter == '_' {
			numberOfLetterMissing++
		}
	}
	if numberOfLetterMissing == 0 {
		isRunning = false
	}
	return isRunning
}
