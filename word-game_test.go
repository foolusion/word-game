package wordgame

import (
  "testing"
)

var checkGuessTests = []struct{
  word, guess string
  ecl, ecp int
  ee error
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
