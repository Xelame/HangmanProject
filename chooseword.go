package main

import (
	"math/rand"
	"time"
)

// Main Function to choose a word randomly in my dictionary
func ChooseWord(dictionary []byte) string {
	dictLenght := CalcNumberOfWord(dictionary)     // Fund number of word
	randomNumber := ChooseRandomNumber(dictLenght) // Generate random number
	return ReadWord(dictionary, randomNumber)      // Return my word
}

// Function read the word in my dictionary
func ReadWord(dictionary []byte, wordPositionChoosen int) string {
	count := 0
	myWord := []byte{}
	for _, bytes := range dictionary {
		if count == wordPositionChoosen-1 { // To write the word I start after the previous line break
			myWord = append(myWord, bytes)
		}
		if bytes == '\n' {
			count++
		}
	}
	return string(myWord) // Return the word
}

// Function to know number of word in the dictionary
func CalcNumberOfWord(dictionary []byte) int {
	count := 0
	for _, bytes := range dictionary {
		if bytes == '\n' { // each ligne break = new word
			count++
		}
	}
	return count
}

// Function to choose a random number beetween 0 and the last word position
func ChooseRandomNumber(numberOfWords int) int {
	// Init Source type with a seed who is always different because the number is affiliated with the clock
	randomSource := rand.NewSource(time.Now().UnixNano())
	// Init Rand type who have a random number in each type
	randomValue := rand.New(randomSource)
	// Fix a limit for the random number who is my number of word
	return randomValue.Intn(numberOfWords)
}
