/* -----------------------------------------------------------------------------------
 * Auteur : BOURRY Nathan et Alexandre ROLLAND                     Créer le : 21/10/21
 * errorDetection.go
 * Programme qui gère la gestion des erreur                             Version : v1.0
 * ---------------------------------------------------------------------------------*/

package main

// -----------------------------------------------------------------------------------
// Partie importation librairie
// -----------------------------------------------------------------------------------

import (
	"log"
	"os"
)

//-------------------------------------------------------------------------------------
//Partie déclaration des constantes
//-------------------------------------------------------------------------------------

const DOCUMENT_ERROR = "errorGestionary.txt"

// -----------------------------------------------------------------------------------
// Partie du programme
// -----------------------------------------------------------------------------------

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
