package main

// -----------------------------------------------------------------------------------
// Import Part
// -----------------------------------------------------------------------------------

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

// -----------------------------------------------------------------------------------
// Const and Var Part
// -----------------------------------------------------------------------------------

const DOCUMENT_ERROR = "errorGestionary.txt"

// -----------------------------------------------------------------------------------
// Program Part
// -----------------------------------------------------------------------------------

func main() {

	Menu()

}

func OpenScanner(fileName string) *bufio.Scanner {
	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		errorDectection(errOpen.Error())
		log.Fatal(TEXT_ERROR_OPEN)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

func errorDectection(errDict string) {
	file, err := os.Create(DOCUMENT_ERROR)
	if err != nil {
		log.Fatal(err)
	} else {
		_, err = file.WriteString(errDict)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
