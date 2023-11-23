package Hangman

import (
	"net/http"
	"strings"
)

func HandleHangman(w http.ResponseWriter, r *http.Request) {
	if RevealedWord == WordToFind {
		http.Redirect(w, r, "/facile", http.StatusSeeOther)
		return
	}
	letter := r.FormValue("letter")
	if letter == "" {
		http.Error(w, "Missing letter parameter", http.StatusBadRequest)
		return
	}

	if !strings.Contains(WordToFind, letter) {
		RemainingAttempts--
	} else {
		for i := 0; i < len(WordToFind); i++ {
			if string(WordToFind[i]) == letter {
				RevealedWord = RevealedWord[:i] + letter + RevealedWord[i+1:]
			}
		}
	}
	http.Redirect(w, r, "/facile", http.StatusSeeOther)
	return
}
