package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Partie recherche du mot dans le fichier
	dictionary, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	wordChoosen := ChooseWord(dictionary) // FIXME Use this
	//Partie José
	PrintJosé(8, "hangman.txt") // TODO Read the file and init a []byte before calling function

	// TODO List of function to show our output (Nathan)
	// TODO Create a hidden word and manipulate it (Alex)

}
