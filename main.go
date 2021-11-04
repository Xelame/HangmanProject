// TODO If it's possible, add language choose and adapt variables 📚
// TODO If it's possible, create document for rules and you can open
// FIXME Problème dans retry ça ne change pas de mots
// FIXME Gestions des constantes (notamment les strings) 🕵️‍♂️
// FIXME Lecture des fichiers (+ entrée joueur) 🧙‍♂️
// FIXME Fractionnage des fichier a revoir
// FIXME Retravailler le desing pour plsu explicite

/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * main.go
 * Programme principale                                                 Version : v1.0
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

const HANGMAN_FILENAME = "hangman.txt"
const DICTIONARY_FILENAME = "words.txt"
const HANGMAN_BANNER = "\n \n \n \n \n \n \033[31m██░ ██  ▄▄▄       ███▄    █   ▄████  ███▄ ▄███▓ ▄▄▄       ███▄    █\n▓██░ ██▒▒████▄     ██ ▀█   █  ██▒ ▀█▒▓██▒▀█▀ ██▒▒████▄     ██ ▀█   █\n▒██▀▀██░▒██  ▀█▄  ▓██  ▀█ ██▒▒██░▄▄▄░▓██    ▓██░▒██  ▀█▄  ▓██  ▀█ ██▒\n░▓█ ░██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒░▓█  ██▓▒██    ▒██ ░██▄▄▄▄██ ▓██▒  ▐▌██▒\n░▓█▒░██▓ ▓█   ▓██▒▒██░   ▓██░░▒▓███▀▒▒██▒   ░██▒ ▓█   ▓██▒▒██░   ▓██░\n▒ ░░▒░▒ ▒▒   ▓▒█░░ ▒░   ▒ ▒  ░▒   ▒ ░ ▒░   ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ \n▒ ░▒░ ░  ▒   ▒▒ ░░ ░░   ░ ▒░  ░   ░ ░  ░      ░  ▒   ▒▒ ░░ ░░   ░ ▒░\n░  ░░ ░  ░   ▒      ░   ░ ░ ░ ░   ░ ░      ░     ░   ▒      ░   ░ ░ \n░  ░  ░      ░  ░         ░       ░        ░         ░  ░         ░\033[0m \n \n \n \n \n \n"
const TEXT_ERROR_DICT = "\n||HOO...no ... José couldn't choose a word :(                   ||\n||Please try to close and open the app again so José can decide!||"
const TEXT_ERROR_NO_WORD = "No Word in file"
const TEXT_ERROR_HANG = "\n||HOO no ... José did not find his rope!                        ||\n||Please close the program and open it so that José can find it.||"
const TEXT_ERROR_NO_CONTENT = "No content in hangman file"
const TEXT_INTRO = "||Welcome to the Hangman game !                           ||\n||Will you be able to find the hidden word and save José ?||"
const TEXT_FINISH_WIN = "Well Played you found the word and save Jose !\nDo you want to retry ? [Y]es or [N]o"
const TEXT_FINISH_LOST = "Poor Jose ...\nRetry your chance for him to survive ? [Y]es or [N]o"
const ATTEMPTS_NUMBER = 4

var input string = ""
var solution = []rune{'-', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'À', 'Á', 'Â', 'Ã', 'Ä', 'Å', 'Ç', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï'}

// -----------------------------------------------------------------------------------
// Partie déclaration des variables
// -----------------------------------------------------------------------------------

var wordChoosen string = ChooseWord(DICTIONARY_FILENAME)
var startHint byte = wordChoosen[len(wordChoosen)/2-1]
var lettersAlreadyAppeard = []rune{'-', ToUpper(rune(startHint))}
var hiddenWord string = HideWord(wordChoosen, &lettersAlreadyAppeard)

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func main() {

	Game(ATTEMPTS_NUMBER)
	Retry()

}
