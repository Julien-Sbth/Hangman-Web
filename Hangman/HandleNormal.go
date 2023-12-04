package Hangman

import (
	"Hangman-Web/Connexion"
	"Hangman-Web/Fonctions"
	"database/sql"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
)

func HandleNormale(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var data Data

	if WordToFind == "" {
		HandleGetWord(w, r, "normal")
	}

	CurrentLevel = "normal"

	if CurrentLevelData != &NormalData {
		NormalData = LevelData{}
	}
	CurrentLevelData = &NormalData

	if CurrentLevelData != &FacileData {
		FacileData = LevelData{}
	}
	CurrentLevelData = &FacileData

	session, err := Connexion.Store.Get(r, Connexion.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	alphabet := make([]string, 0)
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	if username, ok := session.Values["username"].(string); ok {
		data = struct {
			Username           string
			Word               string
			RemainingAtt       int
			Message            string
			Level              string
			GameLost           bool
			GameWon            bool
			Score              int
			Alphabet           []string
			ShowGuessedLetters bool
		}{
			Username:           username,
			Word:               RevealedWord,
			RemainingAtt:       RemainingAttempts,
			Message:            "",
			Level:              "facile",
			GameLost:           false,
			GameWon:            true,
			Score:              Score,
			Alphabet:           alphabet,
			ShowGuessedLetters: true,
		}
	}

	if r.Method == "POST" {
		r.ParseForm()
		level := r.Form.Get("level")

		if level != CurrentLevel {
			Fonctions.ResetGameData(level)
			CurrentLevel = level

			switch level {
			case "facile":
				CurrentLevelData = &FacileData
			case "normal":
				CurrentLevelData = &NormalData
			case "difficile":
				CurrentLevelData = &DifficileData
			}

			ResetAllGameData()
		}
		r.ParseForm()
		if r.URL.Path == "/normal" {
			r.ParseForm()
			guessedLetter := r.Form.Get("letter")
			HandleGuessedLetter(guessedLetter)
		}

		http.Redirect(w, r, "/normal", http.StatusSeeOther)
		return
	}

	if RemainingAttempts == 0 {
		data.GameLost = true
		data.Word = WordToFind
		data.Message = "Vous avez perdu :-("
		RevealedWord = ""
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "normal")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	if RevealedWord == WordToFind {
		data.Message = "Félicitations, vous avez trouvé le mot !"
		Score++
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "normal")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	tmpl, err := template.ParseFiles("templates/html/Level/normal.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
