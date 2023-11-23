package Hangman

import (
	"math/rand"
	"net/http"
	"strings"
)

var score int

func HandleReplay(w http.ResponseWriter, r *http.Request) {
	if CurrentLevel == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	RemainingAttempts = 8
	if CurrentLevel == "normal" {
		RemainingAttempts = 6
	} else if CurrentLevel == "difficile" {
		RemainingAttempts = 4
	}
	// Choose a new random word from the list
	WordToFind = Words[rand.Intn(len(Words))]
	// Reset revealed word
	RevealedWord = ""
	RevealedWord = strings.Repeat("*", len(WordToFind))

	// Réinitialisation du score ici
	Score = 0

	http.Redirect(w, r, "/"+CurrentLevel, http.StatusSeeOther)
}
