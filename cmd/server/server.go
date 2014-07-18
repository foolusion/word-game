package main

import "github.com/foolusion/wordgame"

func main() {
	g, err := NewGame("bust")
	if err != nil {
		panic(err)
	}
}
