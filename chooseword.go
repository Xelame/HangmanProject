package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func ChooseWord() string {
	dictionary, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	dictLenght := CalcNumberOfWord(dictionary)
	randomNumber := ChooseRandomNumber(dictLenght)
	return ReadWord(dictionary, randomNumber)
}

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
