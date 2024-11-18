package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko"
)

func main() {
	game := miko.NewGame()
	ebiten.SetWindowSize(600, 800)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
