package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Main Function to choose a word randomly in my dictionary
func ChooseWord(dictionary string) string {
	var contentOfDictionary = []string{}
	dictionaryFile, errOpen := os.Open(dictionary)
	if errOpen != nil {
		errorDectection(errOpen.Error())
		log.Fatal(TEXT_ERROR_DICT)
	}
	defer dictionaryFile.Close()

	scanner := bufio.NewScanner(dictionaryFile)
	for scanner.Scan() {
		contentOfDictionary = append(contentOfDictionary, scanner.Text())
	}
	fmt.Println(len(contentOfDictionary))
	//randomNumber := ChooseRandomNumber(len(contentOfDictionary)) // Generate random number
	return ReadWord(contentOfDictionary, 6) // Return my word
}

// Function read the word in my dictionary
func ReadWord(fileContent []string, wordPositionChoosen int) string {
	//Partie José initialisation
	var myWord string

	for index, fileWord := range fileContent {
		if index == wordPositionChoosen {
			myWord = fileWord
		}
	}
	return myWord // Return the word
}

// Function to choose a random number beetween 0 and the last word position
func ChooseRandomNumber(numberOfWords int) int {
	// Init Source type with a seed who is always different because the number is affiliated with the clock
	randomSource := rand.NewSource(time.Now().UnixNano())
	// Init Rand type who have a random number in each type
	randomValue := rand.New(randomSource)
	// Fix a limit for the random number who is my number of word
	return randomValue.Intn(numberOfWords + 1)
}
