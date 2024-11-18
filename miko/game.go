package miko

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ysfgrl/miko/miko/assets"
	"github.com/ysfgrl/miko/miko/audio"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/input"
	"github.com/ysfgrl/miko/miko/scene"
	"image"
	"math/rand"
	"time"
)

var bg *ebiten.Image

func init() {
	rand.Seed(time.Now().UnixNano())

	f, err := assets.Assets.Open("images/bg/2.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	bg = ebiten.NewImageFromImage(img)
}

type Game struct {
	scene core.Scene
}

func NewGame() *Game {

	game := &Game{
		scene: &scene.Start{},
	}
	go func() {
		if err := audio.Load(); err != nil {
			panic(err)
		}
	}()
	if err := audio.Finalize(); err != nil {
		panic(err)
	}
	return game
}

func (g *Game) Update() error {
	input.Current().Update()
	g.scene = g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(bg, op)
	g.scene.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %.2f", ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 800
}
