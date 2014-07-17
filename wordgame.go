package wordgame

type game struct {
	word       string
	numGuesses int
}

func NewGame(word string) *game {
  // check word for repeated letters.
  return &game{
    word: word,
    numGuesses: 0,
  }
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

	gl, err := letterMap(guess)
	if err != nil {
	  return 0, 0, err
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

func letterMap(guess string) (map[rune]int, error) {
  gl := make(map[rune]int, len(guess))
  for i, l := range guess {
		if _, ok := gl[l]; ok {
			return nil, ErrRepeatedLetter(guess)
		} else {
			gl[l] = i
		}
	}
	return gl, nil
}