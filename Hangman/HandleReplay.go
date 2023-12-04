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
	WordToFind = Words[rand.Intn(len(Words))]
	RevealedWord = ""
	RevealedWord = strings.Repeat("*", len(WordToFind))

	Score = 0

	http.Redirect(w, r, "/"+CurrentLevel, http.StatusSeeOther)
}
