package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/audio"
	"github.com/ysfgrl/miko/miko/core"
)

var assets = []*ebiten.Image{}

func init() {
	assets, _ = core.LoadAssets("images/life")
}

type Life struct {
	position core.Position
	life     int
	life2    int
}

func NewLife(life2 int) *Life {
	return &Life{
		position: core.Position{
			X: 0,
			Y: 0,
		},
		life2: life2,
		life:  6,
	}
}

func (l *Life) IsDead() bool {
	return l.life == 0
}

func (l *Life) Damage() int {
	return l.life2
}
func (l *Life) SetLife(life int) {
	l.life += life
	if life > 0 {
		audio.PlaySE(audio.SE_ITEMGET)
	} else {
		audio.PlaySE(audio.SE_DAMAGE)
	}
	if l.life <= 0 {
		l.life = 0
	}
	if l.life >= len(assets) {
		l.life = len(assets) - 1
	}

}
func (l *Life) Dec() {
	l.life -= 1
	if l.life <= 0 {
		l.life = 0
	}
	audio.PlaySE(audio.SE_DAMAGE)
}

func (l *Life) Inc() {
	l.life += 1
	if l.life >= len(assets) {
		l.life = len(assets) - 1
	}
	audio.PlaySE(audio.SE_ITEMGET)
}

func (l *Life) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(assets[l.life], op)
}
