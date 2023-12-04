package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func HandleGetWord(w http.ResponseWriter, r *http.Request, difficulty string) {
	var filePath string

	switch difficulty {
	case "facile":
		filePath = "templates/words/facile_words.txt"
	case "normal":
		filePath = "templates/words/normal_words.txt"
	case "difficile":
		filePath = "templates/words/difficile_words.txt"
	default:
		filePath = "templates/words/default_words.txt"
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Words = append(Words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	WordToFind = Words[rand.Intn(len(Words))]
	RevealedWord = strings.Repeat("*", len(WordToFind))
	RemainingAttempts = 6

}
