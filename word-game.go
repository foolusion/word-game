package wordgame

type game struct {
	word       string
	numGuesses int
}

type ErrGuessLen string

func (e ErrGuessLen) Error() string {
	return "ErrGuessLen: " + string(e) + " must be the same length as the answer."
}

type ErrRepeatedLetter string

func (e ErrRepeatedLetter) Error() string {
	return "ErrRepeatedLetter: " + string(e) + " cannot use the same letter twice."
}

func checkGuess(word, guess string) (int, int, error) {
	if len(guess) != len(word) {
		return 0, 0, ErrGuessLen(guess)
	}

	gl := make(map[rune]int, 4)

	for i, l := range guess {
		if _, ok := gl[l]; ok {
			return 0, 0, ErrRepeatedLetter(guess)
		} else {
			gl[l] = i
		}
	}

	correctLetters, correctPositions := 0, 0
	for i, l := range word {
		if v, ok := gl[l]; ok {
			correctLetters++
			if v == i {
				correctPositions++
			}
		}
	}

	return correctLetters, correctPositions, nil
}
