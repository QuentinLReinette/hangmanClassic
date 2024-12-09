package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"strings"
)

// structure qui contient les variables de la partie
type data struct {
	Mot    string      // mot à trouver
	Vies   int         // nombre de vies restantes
	Out    []rune      // état actuel du mot avec des underscores
	Essais [2][][]rune // liste des essais précédents : [0] lettres, [1] mots
	Ascii  string      // alphabet ascii art à utiliser pour l'affichage de Out
}

var Game data // structure qui contient les variables de la partie

// affiche message de défaite
func defaite() {
	Prntout()
	os.Exit(0)
}

// affiche message de victoire
func victoire() {
	fmt.Println("Vous avez trouvé le mot, félicitations !")
	os.Exit(0)
}

// affiche la partie
func Prntout() {
	clearT()
	pendu()
	if Game.Vies == 0 {
		fmt.Printf("Vous avez perdu. Le mot était : %s\n", Game.Mot)
		return
	}
	if Game.Ascii == "" {
		Game.Ascii = "standard.txt"
	}
	ascii()
	fmt.Println()
}

// affiche Game.Out en ascii art
func ascii() {
	alphB := Openfile(Game.Ascii) // dictionnaire ascii art
	for i := 1; i <= 8; i++ {     // affiche Game.Out ligne par ligne
		for k, j := range Game.Out {
			l := (9 * (int(j - 32))) + i // converti la lettre en ligne dans alphB
			fmt.Print(alphB[l][0 : len(alphB[l])-1])
			if k < len(Game.Out)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// affiche le pendu
func pendu() {
	dessins := Openfile("hangman.txt")
	if Game.Vies == 10 {
		fmt.Print("\n\n\n\n\n\n\n\n")
	} else {
		num := 9 - Game.Vies
		fmt.Println()
		for i := (8 * num); i < (8*num)+7; i++ {
			fmt.Println(dessins[i][0:9])
		}
	}
}

// fonction principale.
// vérifie si la lettre est dans le mot
func InputCheck() {
	var in string

	fmt.Print("Essais : ")
	for i, a := range Game.Essais { // affiche la liste des lettres/mots essayés
		for j, b := range a {
			fmt.Print(string(b))
			if j < len(a)-1 {
				fmt.Print(", ")
			}
		}
		if i < 1 && len(Game.Essais[1]) > 0 {
			fmt.Print(" - ")
		}
	}
	fmt.Println()

	fmt.Print("\nLettre / Mot : ")
	fmt.Scanf("%s", &in)
	in = strings.ToLower(in) // passe les Maj en Min

	if in == "stop" { // sauvegarde et quitte
		save()
		os.Exit(0)
	}

	if in == "" {
		Prntout()
		InputCheck()
		return
	}

	for _, i := range in { // vérifie si que des lettres
		if !(i >= 'a' && i <= 'z') && !(i >= 'A' && i <= 'Z') {
			Prntout()
			InputCheck()
			return
		}
	}

	Tin := []rune(in)
	for _, i := range Game.Essais { // compare avec liste des mots et lettres déjà essayés
		for _, j := range i {
			if string(Tin) == string(j) {
				Prntout()
				InputCheck()
				return
			}
		}
	}

	if len(Tin) == 1 { // cherche la lettre dans le mot
		Game.Essais[0] = append(Game.Essais[0], Tin) // ajoute à la liste des lettres déjà essayées

		C := 0
		for i, j := range Game.Mot { // révele et compte toutes les instances de la lettre dans le mot
			if j == Tin[0] {
				Game.Out[i] = j
				C++
			}
		}
		if C == 0 {
			Game.Vies--
		}
		Prntout()
		if string(Game.Out) == Game.Mot {
			victoire()
		}
	} else { // compare le mot entré avec celui à deviner
		Game.Essais[1] = append(Game.Essais[1], Tin) // ajoute à la liste des mots déjà essayés

		if string(Tin) == Game.Mot {
			victoire()
		} else {
			Game.Vies -= 2
			if Game.Vies > 0 {
				Prntout()
			}
		}
	}

	if Game.Vies <= 0 {
		Game.Vies = 0
		defaite()
	} else {
		InputCheck()
	}
}

// ouvre un fichier
func Openfile(chemin string) (Liste []string) {
	var readFile *os.File

	readFile, err1 := os.Open(chemin) // ouvre fichier
	if err1 != nil {
		b, err2 := os.Open("../" + chemin)
		if err2 != nil {
			log.Fatalln(err1)
		}
		readFile = b
	}

	fileScanner := bufio.NewScanner(readFile) // lit fichier

	fileScanner.Split(bufio.ScanLines) // formate liste

	for fileScanner.Scan() { // copie liste dans []string
		Liste = append(Liste, fileScanner.Text())
	}

	readFile.Close() // ferme liste

	return
}

// demande de choisir la difficulté
func ChoixDiff() string {
	var diff string
	fmt.Println("Choisissez la difficulté : 1, 2, 3")
	fmt.Scanf("%s", &diff)

	if diff != "1" && diff != "2" && diff != "3" {
		fmt.Println("Difficulté incorrecte. Le choix par défaut est 1. Entrée pour continuer")
		fmt.Scanf("%s")
		return "words1.txt"
	}
	if diff == "1" {
		return "words1.txt"
	} else if diff == "2" {
		return "words2.txt"
	} else {
		return "words3.txt"
	}
}

// ajoute des lettres aléatoires dans Game.Out
func RdmLttrs() {
	var cont int
	var dejafait bool

	for i := 0; i <= (len(Game.Mot)/2)-1; i++ {
		lettre := rune(Game.Mot[rand.IntN(len(Game.Mot)-1)]) // sélectionne une lettre aléatoire
		cont = 0
		dejafait = false

		for _, a := range Game.Out { // vérifie si lettre déjà dans Game.Out
			if a == lettre {
				dejafait = true
			}
		}
		for _, j := range Game.Mot { // compte les occurences
			if j == lettre {
				cont++
			}
		}
		if dejafait || i+cont-1 > (len(Game.Mot)/2)-1 {
			i--
			continue
		}
		for k, l := range Game.Mot { // affiche lettre dans Game.Out
			if l == lettre {
				Game.Out[k] = lettre
			}
		}

		i += cont
		Game.Essais[0] = append(Game.Essais[0], []rune{lettre})
	}
}

func clearT() { // efface le terminal
	var cmd *exec.Cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// charge une partie sauvegardée
func Load(savefile string) {
	partie := Openfile(savefile) // lit le fichier de sauvegarde
	if len(partie) == 0 {
		log.Fatal("Fichier vide.")
	}
	err := json.Unmarshal([]byte(partie[0]), &Game) // récupère les valeurs de Game
	if err != nil {
		log.Fatal(err)
	}
	pendu()
}

// sauvegarde la partie
func save() {
	partie, err1 := json.Marshal(Game) // convertit Game en JSON
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := os.WriteFile("partie.txt", partie, 0644) // crée ou écrase partie.txt avec partie
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Partie sauvegardée dans partie.txt. Lancez la prochaine partie avec \"--startWith partie.txt\" pour continuer.")
}
