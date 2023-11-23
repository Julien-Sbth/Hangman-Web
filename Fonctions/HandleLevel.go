package Fonctions

import (
	"Hangman-Web/Connexion"
	"database/sql"
	"html/template"
	"net/http"
)

type LevelData struct {
	WordToFind        string
	RevealedWord      string
	RemainingAttempts int
}

var CurrentLevel = ""
var LevelDataMap = make(map[string]LevelData)

func ResetGameData(level string) {
	// Réinitialiser les données spécifiques à un niveau
	LevelDataMap[level] = LevelData{
		WordToFind:        "", // Réinitialisez avec les valeurs appropriées
		RevealedWord:      "", // Réinitialisez avec les valeurs appropriées
		RemainingAttempts: 0,  // Réinitialisez avec les valeurs appropriées
	}
}

func HandleLevel(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	session, err := Connexion.Store.Get(r, Connexion.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if username, ok := session.Values["username"].(string); ok {
		tmpl, err := template.ParseFiles("templates/html/Level/level.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Username string
		}{
			Username: username,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		level := r.Form.Get("level")

		// Réinitialiser les données lorsque l'utilisateur change de niveau
		if level != CurrentLevel {
			ResetGameData(level) // Réinitialiser les données spécifiques au niveau sélectionné
			CurrentLevel = level
		}
	}

	tmpl, err := template.ParseFiles("templates/html/Level/level.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
