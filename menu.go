package main

//-------------------------------------------------------------------------------------
// Import Part
//-------------------------------------------------------------------------------------

import (
	"fmt"
	"os"
	"os/exec"
)

//-------------------------------------------------------------------------------------
// Const and Var Part
//-------------------------------------------------------------------------------------

const ATTEMPTS_NUMBER = 10

const HANGMAN_BANNER = `██░ ██  ▄▄▄       ███▄    █   ▄████  ███▄ ▄███▓ ▄▄▄       ███▄    █
▓██░ ██▒▒████▄     ██ ▀█   █  ██▒ ▀█▒▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █
▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒
░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒
░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░
▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ ░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░
░  ░░ ░  ░   ▒      ░   ░ ░ ░ ░   ░ ░      ░     ░   ▒      ░   ░ ░ 
░  ░  ░      ░  ░         ░       ░        ░         ░  ░         ░`

const TEXT_INTRO = `╔═══════════════════════════════════════════════════════════════════╗
║                                                                   ║
║                              1.Play                               ║
║                              2.Rules                              ║
║                              3.Credits                            ║
║                              4.Quit                               ║
║                                                                   ║
╚═══════════════════════════════════════════════════════════════════╝`

//-------------------------------------------------------------------------------------
// Program Part
//-------------------------------------------------------------------------------------

// Function to display a menu at the beginning
func Menu() {
	Clear()

	// Banner
	fmt.Print("\n \n \033[31m")
	fmt.Print(HANGMAN_BANNER)
	fmt.Print("\033[0m \n \n \n")

	// Introduction
	fmt.Println(TEXT_INTRO + "\n \n")

	// Choose of player
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// Run the game
		Game(ATTEMPTS_NUMBER)
	case "2":
		// Show the rule
		OpenRules("Rules.txt")
	case "3":
		// Show our Name
		Clear()
		fmt.Print("Developped by Nathan Bourry and Alexandre Rolland")
		fmt.Print("\n\nPress [ENTER] to return to the menu")
		fmt.Scanf("%v")
		Menu()
	case "4":
		// Stop the program
		fmt.Println("See you later !")
	default:
		// Anything else reset the menu
		Clear()
		Menu()
	}
}

// Function to read differently an file.txt (full only)
func OpenRules(rulesFileName string) {
	Clear()
	cmd := exec.Command("cat", rulesFileName)
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Print("\n\nPress [ENTER] to return to the menu")
	fmt.Scanf("%v")
	Menu()
}
