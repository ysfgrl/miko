package core

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() Scene
	Draw(screen *ebiten.Image)
}
