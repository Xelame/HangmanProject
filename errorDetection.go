package main

import (
	"log"
	"os"
)

func errorDectection(errDict string) {
	file, err := os.Create("errorGestionary.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		_, err = file.WriteString(errDict)
		if err != nil {
			log.Fatal(err)
		}
	}
}
