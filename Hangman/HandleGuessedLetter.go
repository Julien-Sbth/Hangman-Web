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

		if RevealedWord == WordToFind {
			Score++
		}
	}
}
