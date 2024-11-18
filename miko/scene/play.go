package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/audio"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/entity"
	"github.com/ysfgrl/miko/miko/input"
)

type Play struct {
	Player  *entity.Entity
	Factory *entity.Factory
}

func NewPlay() core.Scene {
	player := entity.NewEntity("images/player/child", &core.Position{
		X: 0,
		Y: 710,
	})
	player.AddAction(input.ActionLeft, func(e *entity.Entity) {
		e.WalkLeft()
	})
	player.AddAction(input.ActionRight, func(e *entity.Entity) {
		e.WalkRight()
	})
	return &Play{
		Player:  player,
		Factory: entity.NewFactory(5),
	}
}

func (p *Play) Update() core.Scene {
	last := p.Factory.Update()
	p.Player.Update()

	for i := range last {
		box1 := last[i].GetBox()
		box2 := p.Player.GetBox()

		if box2.Top <= box1.Bottom {
			if (box1.Left <= box2.Left && box2.Left <= box1.Right) ||
				(box1.Left <= box2.Right && box2.Right <= box1.Right) {
				p.Factory.Pop(i)
				p.Player.Life.SetLife(last[i].Life.Damage())
			}
			if box1.Left >= box2.Left && box1.Right <= box2.Right {
				p.Factory.Pop(i)
				p.Player.Life.SetLife(last[i].Life.Damage())
			}

		}
	}
	if p.Player.Life.IsDead() {
		audio.PauseBGM()
		return &End{}
	}

	return p
}

func (p *Play) Draw(screen *ebiten.Image) {
	p.Factory.Draw(screen)
	p.Player.Draw(screen)
	p.Player.Life.Draw(screen)
}
