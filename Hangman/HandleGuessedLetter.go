package Hangman

func HandleGuessedLetter(letter string) {
	if len(letter) > 0 {
		guessedLetter := []rune(letter)[0]
		found := false
		updatedRevealedWord := ""

		for i, char := range WordToFind {
			if char == guessedLetter {
				updatedRevealedWord += string(guessedLetter)
				found = true
			} else {
				updatedRevealedWord += string([]rune(RevealedWord)[i])
			}
		}

		if !found {
			RemainingAttempts--
		}

		RevealedWord = updatedRevealedWord

		// Vérifie si le mot a été trouvé après chaque tentative de deviner une lettre
		if RevealedWord == WordToFind {
			Score++
		}
	}
}
