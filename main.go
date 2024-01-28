package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lidldev/GopherJump/game"
)

func main() {
	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
