[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/1YLV-els)

# Hangman - Jeu du Pendu

Ce programme est une version du jeu du pendu en Go, où le joueur doit deviner un mot en proposant des lettres ou des mots.

## Fonctionnalités

- Choix de la difficulté : Trois niveaux disponibles avec des fichiers de mots différents.
- Sauvegarde et chargement : La partie peut être sauvegardée et chargée à tout moment.
- ASCII art : Affichage du mot à deviner avec un alphabet ASCII.
- Affichage du pendu : Visualisation du pendu à chaque erreur.

## Utilisation

### Choisir la difficulté
Lancer la partie avec un fichier de mots :
```bash
go run main.go words1.txt
```
Ou sans arguments et le jeu demandera de choisir :
```bash
go run main.go
```

### Sauvegarder et charger une partie
- Sauvegarder en entrant `stop`
- Charger une partie avec :
```bash
go run main.go --startWith partie.txt
```

### Utilisation d'ASCII art
Utiliser un fichier pour l'affichage ASCII art que celui par défaut :
```bash
go run main.go --letterFile [alphabet_ascii]
```
