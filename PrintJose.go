package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintJose(nbrs_tentative int, hangmanFileName string) {

	//Partie José initialisation
	contenuHangmanByte, errOpen := os.Open(hangmanFileName)
	if errOpen != nil {
		errorDectection(errOpen.Error())
		log.Fatal(TEXT_ERROR_HANG)
	}
	defer contenuHangmanByte.Close()

	reader := bufio.NewScanner(contenuHangmanByte)

	for i := 0; reader.Scan(); i++ {
		if (10-nbrs_tentative)*8 <= i+1 && i+1 <= (10-nbrs_tentative+1)*8 { // Constantes !
			fmt.Println(reader.Text())
		}
	}
}
