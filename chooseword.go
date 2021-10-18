package main

import (
	"math/rand"
	"time"
)

// Main Function who return our word in string
func ChooseWord(dictionary []byte) string {
	dictLenght := CalcNumberOfWord(dictionary)
	randomNumber := ChooseRandomNumber(dictLenght)
	return ReadWord(dictionary, randomNumber)
}

//
func ReadWord(dictionary []byte, wordPositionChoosen int) string {
	count := 0
	myWord := []byte{}
	for _, bytes := range dictionary {
		if count == wordPositionChoosen-1 {
			myWord = append(myWord, bytes)
		}
		if bytes == '\n' {
			count++
		}
	}
	return string(myWord[:len(myWord)])
}
func CalcNumberOfWord(dictionary []byte) int {
	count := 0
	for _, bytes := range dictionary {
		if bytes == '\n' {
			count++
		}
	}
	return count
}

func ChooseRandomNumber(numberOfWords int) int {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomValue := rand.New(randomSource)
	return randomValue.Intn(numberOfWords)
}
