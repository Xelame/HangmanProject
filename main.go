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

// Run the program
func main() {

	Menu()

}

// Function to open a file and create a scanner for this file
func OpenScanner(fileName string) *bufio.Scanner {
	file, errOpen := os.Open(fileName) // Open the file
	if errOpen != nil {                // Error detection
		ErrorDectection(errOpen.Error())
	}
	scanner := bufio.NewScanner(file) // Create a scanner
	return scanner
}

// Function to print the type of the error in a document
func ErrorDectection(errFile string) {
	// Create every time the same file to reset at each error dectected
	file, err := os.Create(DOCUMENT_ERROR)
	// Test if the creation have an error
	if err != nil {
		log.Fatal(err)
	} else {
		// Write the type of error detected in this file
		_, err = file.WriteString(errFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Clear the console display
func Clear() {
	cmd := exec.Command("clear") // Assigns to the variable the path to launch the commands + the specified string
	cmd.Stdout = os.Stdout       // Assigns the output value to that of our console
	cmd.Run()                    // Executes the command that was specified as a string
}
