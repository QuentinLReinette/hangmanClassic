package main

import (
	"hangman"
	"math/rand/v2"
	"os"
)

func main() {
	var Liste []string
	var chemin string

	if len(os.Args) == 3 && os.Args[1] == "--letterFile" {
		hangman.Game.Ascii = os.Args[2]
	}
	if len(os.Args) == 3 && os.Args[1] == "--startWith" {
		hangman.Load(os.Args[2])
	} else {
		if len(os.Args) == 2 {
			chemin = os.Args[1]
		} else {
			chemin = hangman.ChoixDiff() // choix de difficulté
		}
		Liste = hangman.Openfile(chemin)

		hangman.Game.Mot = Liste[rand.IntN(len(Liste)-1)] // prend un mot au hasard

		for range hangman.Game.Mot {
			hangman.Game.Out = append(hangman.Game.Out, '_')
		}
		hangman.RdmLttrs()
		hangman.Game.Vies = 10
	}

	hangman.Prntout()
	hangman.InputCheck()
}
