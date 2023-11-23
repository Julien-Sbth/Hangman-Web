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

func HandleDifficile(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var data Data

	CurrentLevel = "difficile"

	// Lorsque vous changez de niveau, réinitialisez le jeu spécifique à ce niveau
	if CurrentLevelData != &DifficileData {
		DifficileData = LevelData{} // Réinitialisation des données du niveau facile
	}
	CurrentLevelData = &DifficileData // Mettez à jour les données du niveau actuel

	if WordToFind == "" {
		HandleGetWord(w, r, "difficile")
	}

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

		// Réinitialiser les données lorsque l'utilisateur change de niveau
		if level != CurrentLevel {
			// Réinitialiser les données spécifiques au niveau sélectionné
			Fonctions.ResetGameData(level)
			CurrentLevel = level

			// Réinitialiser le jeu spécifique au niveau actuel
			switch level {
			case "facile":
				CurrentLevelData = &FacileData
			case "normal":
				CurrentLevelData = &NormalData
			case "difficile":
				CurrentLevelData = &DifficileData
			}

			ResetAllGameData() // Effacer toutes les données spécifiques à tous les niveaux
		}
		r.ParseForm()
		if r.URL.Path == "/difficile" {
			r.ParseForm()
			guessedLetter := r.Form.Get("letter")
			HandleGuessedLetter(guessedLetter)
		}

		http.Redirect(w, r, "/difficile", http.StatusSeeOther)
		return
	}

	if RemainingAttempts == 0 {
		data.GameLost = true
		data.Word = WordToFind
		data.Message = "Vous avez perdu :-("
		RevealedWord = ""
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "difficile")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	if RevealedWord == WordToFind {
		data.Message = "Félicitations, vous avez trouvé le mot !"
		Score++
		RemainingAttempts = 10
		if len(Words) == 0 {
			HandleGetWord(w, r, "difficile")
		}
		WordToFind = Words[rand.Intn(len(Words))]
		RevealedWord = strings.Repeat("*", len(WordToFind))
	}

	tmpl, err := template.ParseFiles("templates/html/Level/difficile.html")
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
