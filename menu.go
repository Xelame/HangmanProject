package main

import "fmt"

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
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
║                                                                   ║
╚═══════════════════════════════════════════════════════════════════╝`

func Menu() {
	Clear()

	// Banner
	fmt.Print("\n \n \033[31m")
	fmt.Print(HANGMAN_BANNER)
	fmt.Print("\033[0m \n \n \n")

	// Introduction
	fmt.Println(TEXT_INTRO + "\n \n")

	fmt.Scanf("%s", &input)

	switch input {
	case "1":
		Game(ATTEMPTS_NUMBER)
	case "2":
		Clear()
		fmt.Println("Rule")
	case "3":
		Clear()
		fmt.Println("Created by Alexandre Rolland and Nathan Bourry")
	}

}
