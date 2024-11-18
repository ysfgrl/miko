package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/audio"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/input"
	"image/color"
)

type Start struct {
	timer int
}

func (t *Start) Update() core.Scene {
	t.timer++

	if (input.Current().IsActionKeyJustPressed() || input.Current().IsSpaceJustTouched()) && t.timer > 5 {
		if err := audio.PlayBGM(audio.BGM1); err != nil {
			panic(err)
		}
		return NewPlay()
	}
	ebiten.SetWindowTitle("Miko ve Medo")
	return t
}

func (t *Start) Draw(screen *ebiten.Image) {
	clr := color.Black
	str := "Başlamak İçin Enter a Basın"
	x := (600 - core.Width(str)) / 2
	y := (800-240)/2 + 160
	core.DrawText(screen, str, x, y, clr)
}
