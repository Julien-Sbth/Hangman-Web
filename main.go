package main

import (
	"Hangman-Web/Connexion"
	"Hangman-Web/Fonctions"
	"Hangman-Web/Hangman"
	"fmt"
	"log"
	"net/http"
)

func main() {

	certFile := "KeyHTTPS/certificat.crt"
	keyFile := "KeyHTTPS/privatekey.key"

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/"))))
	http.HandleFunc("/", Fonctions.HandlePlay)
	http.HandleFunc("/facile", Hangman.HandleFacile)
	http.HandleFunc("/normal", Hangman.HandleNormale)
	http.HandleFunc("/difficile", Hangman.HandleDifficile)
	http.HandleFunc("/level", Fonctions.HandleLevel)
	http.HandleFunc("/replay", Hangman.HandleReplay)
	http.HandleFunc("/hangman", Hangman.HandleHangman)
	http.HandleFunc("/connexion", Connexion.HandleConnexion)
	http.HandleFunc("/inscription", Connexion.HandleInscription)
	http.HandleFunc("/reset", Hangman.HandleReset)
	http.HandleFunc("/password", Connexion.HandleForgetPassword)

	fmt.Println("Server started on port :443 https://localhost:443")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("Erreur de démarrage du serveur HTTPS : ", err)
	}
}
