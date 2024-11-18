package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/audio"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/input"
	"golang.org/x/text/language"
	"image/color"
)

type End struct {
	timer          int
	bgmFadingTimer int
	state          int
	texts          map[language.Tag][]string
}

func (e *End) Update() core.Scene {

	if input.Current().IsActionKeyJustPressed() || input.Current().IsSpaceJustTouched() {
		if err := audio.PlayBGM(audio.BGM1); err != nil {
			panic(err)
		}
		return NewPlay()
	}
	return e
}

func (e *End) Draw(screen *ebiten.Image) {
	clr := color.Black
	str := "Başlamak İçin Entera Basın"
	x := (600 - core.Width(str)) / 2
	y := (800-240)/2 + 160
	core.DrawText(screen, str, x, y, clr)
}
