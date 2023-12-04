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

type Data struct {
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
}
type LevelData struct {
	WordToFind        string
	RevealedWord      string
	RemainingAttempts int
}

var CurrentLevel = ""
var LevelDataMap = make(map[string]LevelData)
var Words []string
var WordToFind = ""
var RevealedWord = ""
var RemainingAttempts int
var Score int
var alphabet []string

var FacileData = LevelData{}
var NormalData = LevelData{}
var DifficileData = LevelData{}
var CurrentLevelData *LevelData

func HandleFacile(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var data Data

	if WordToFind == "" {
		HandleGetWord(w, r, "facile")
	}

	CurrentLevel = "facile"

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

		if level != CurrentLevel {
			Fonctions.ResetGameData(level)
			CurrentLevel = level
		}
		if r.URL.Path == "/facile" {
			r.ParseForm()
			guessedLetter := r.Form.Get("letter")
			HandleGuessedLetter(guessedLetter)
		}

		http.Redirect(w, r, "/facile", http.StatusSeeOther)
		return
	}

	if RemainingAttempts == 0 {
		data.GameLost = true
		data.Word = WordToFind
		data.Message = "Vous avez perdu :-("
		RevealedWord = ""
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "facile")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	if RevealedWord == WordToFind {
		data.Message = "Félicitations, vous avez trouvé le mot !"
		Score++
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "facile")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	tmpl, err := template.ParseFiles("templates/html/Level/facile.html")
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

func ResetAllGameData() {
	LevelDataMap = make(map[string]LevelData)
	FacileData = LevelData{}
	NormalData = LevelData{}
	DifficileData = LevelData{}
	CurrentLevel = ""
	Words = []string{}
	WordToFind = ""
	RevealedWord = ""
	RemainingAttempts = 0
	Score = 0
}

func HandleReset(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ResetAllGameData()
		http.Redirect(w, r, "/level", http.StatusSeeOther)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
