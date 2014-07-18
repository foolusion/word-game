package main

import "github.com/foolusion/wordgame"

func main() {
	g, err := wordgame.NewGame("bust")
	if err != nil {
		panic(err)
	}
	g.Main()
}
