package main

import (
	"log"
	"os"
)

func errorDectection(errDict error) {
	file, err := os.Create("errorGestionary.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err = file.WriteString(errDict.Error())
		if err != nil {
			log.Fatal(err)
		}
	}
}
