// TODO If it's possible, add language choose and adapt variables 📚
// TODO If it's possible, create document for rules and you can open
// FIXME Problème dans retry ça ne change pas de mots
// FIXME Lecture des fichiers (+ entrée joueur) 🧙‍♂️
// FIXME Fractionnage des fichier a revoir
// FIXME Retravailler le desing pour plsu explicite

//Commentaire, menu, règle, langue

/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 18/10/21
 * main.go
 * Programme principale                                                 Version : v1.0
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

func main() {
	openRules("Rules.txt")
	Game(ATTEMPTS_NUMBER)
	Retry()

}
