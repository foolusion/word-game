package wordgame

import (
	"fmt"
)

type gameState int

const (
	notStarted gameState = iota
	started
	finished
)

// Game is the main data structure for games.
type Game struct {
	word       string
	numGuesses int
	state      gameState
}

// NewGame creates a new game or returns an error if the word is illegal.
func NewGame(word string) (*Game, error) {
	_, err := letterMap(word)
	if err != nil {
		return nil, err
	}
	return &Game{
		word:       word,
		numGuesses: 0,
		state:      notStarted,
	}, nil
}

// Main is the main game loop.
func (g *Game) Main() {
	for g.state != finished {
		// handle input
		var guess string
		fmt.Scanln(&guess)
		// perform game logic
		cl, cp, err := checkGuess(g.word, guess)
		if err != nil {
			fmt.Println(err)
		} else {
			g.numGuesses++
		}
		if cp == len(g.word) {
			g.state = finished
		}
		// render
		if g.state == finished {
			fmt.Println("You guessed it!")
			return
		}
		fmt.Println(cl, cp)
		fmt.Println("Guess again.")
	}
}

// ErrGuessLen occurs when the length of a guess does not match the length of the game word.
type ErrGuessLen string

func (e ErrGuessLen) Error() string {
	return "ErrGuessLen: " + string(e) + " must be the same length as the answer."
}

// ErrRepeatedLetter occurs when the game word or guess has repeated letters.
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
		}
		gl[l] = i
	}
	return gl, nil
}
