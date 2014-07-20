package wordgame

import (
	"testing"
)

var checkGuessTests = []struct {
	word, guess string
	ecl, ecp    int
	ee          error
}{
	{"bird", "bird", 4, 4, nil},
	{"func", "test", 0, 0, ErrRepeatedLetter("test")},
	{"func", "longword", 0, 0, ErrGuessLen("longword")},
	{"abcdefghijklmnopqrstuvwxyz", "shortword", 0, 0, ErrGuessLen("shortword")},
	{"abcd", "1234", 0, 0, nil},
	{"abcd", "aefg", 1, 1, nil},
	{"axyz", "bacd", 1, 0, nil},
	{"abcd", "dcba", 4, 0, nil},
	{"ycab", "cdab", 3, 2, nil},
}

func TestCheckGuess(t *testing.T) {
	for _, tt := range checkGuessTests {
		acl, acp, ae := checkGuess(tt.word, tt.guess)
		if acl != tt.ecl || acp != tt.ecp || ae != tt.ee {
			t.Errorf("checkGuess(%v, %v): expected %d %d %v, actual %d %d %v", tt.word, tt.guess, tt.ecl, tt.ecp, tt.ee, acl, acp, ae)
		}
	}
}

var newGameTests = []struct {
	word   string
	eGame  *Game
	eError error
}{
	{"bird", &Game{word: "bird", numGuesses: 0, state: notStarted}, nil},
	{"aa", nil, ErrRepeatedLetter("aa")},
	{"Fart", &Game{word: "Fart", numGuesses: 0, state: notStarted}, nil},
	{"abcd", &Game{word: "abcd", numGuesses: 0, state: notStarted}, nil},
}

func TestNewGame(t *testing.T) {
	for _, tt := range newGameTests {
		aGame, aError := NewGame(tt.word)
		if (aGame == nil && tt.eGame != nil) || (aGame != nil && tt.eGame == nil) {
			t.Errorf("NewGame(%v): expected %+v %v, actual %+v %v", tt.word, tt.eGame, tt.eError, aGame, aError)
			continue
		}
		// check if actual and expected are nil and errors are not the same.
		if aGame == nil && tt.eGame == nil && aError != tt.eError {
			t.Errorf("NewGame(%v): expected %+v %v, actual %+v %v", tt.word, tt.eGame, tt.eError, aGame, aError)
			continue
		} else if aGame == nil && tt.eGame == nil {
			// if both are nil and error is same skip checking values
			continue
		}
		// check if the values are the same
		if aGame.word != tt.eGame.word || aGame.numGuesses != tt.eGame.numGuesses || aGame.state != tt.eGame.state {
			t.Errorf("NewGame(%v): expected %+v %v, actual %+v %v", tt.word, tt.eGame, tt.eError, aGame, aError)
		}
	}
}
