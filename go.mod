module Hangman-Web

replace Fonction/Fonction => ./Fonction

replace Fonction/Connexion => ./Connexion

replace Fonction/Level => ./Level

replace Fonction/Hangman => ./Hangman

require (
	github.com/gorilla/sessions v1.2.2
	github.com/mattn/go-sqlite3 v1.14.18
	golang.org/x/crypto v0.15.0
)

require github.com/gorilla/securecookie v1.1.2 // indirect

go 1.21
