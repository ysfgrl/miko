package entity

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/input"
	"math/rand"
)

type Factory struct {
	Entity []core.Queue[*Entity]
}

func NewFactory(size int) *Factory {
	return &Factory{
		Entity: make([]core.Queue[*Entity], size),
	}
}

func (f *Factory) Update() []*Entity {

	for i := range f.Entity {
		if f.Entity[i].Length() == 0 || (f.Entity[i].Length() > 0 && f.Entity[i].Last().GetPosition().Y > 128) {
			newEntity := NewEntity(fmt.Sprintf("images/medo/%d", rand.Intn(6)+1), &core.Position{
				Y: 0,
				X: float64(i * 120),
			})
			newEntity.AddAction(input.ActionRefresh, func(entity *Entity) {
				entity.WalkDown()
			})
			f.Entity[i].Push(newEntity)
		}
	}

	for i := range f.Entity {
		f.Entity[i].For(func(item *Entity) {
			item.Update()
		})
	}

	for i := range f.Entity {
		if f.Entity[i].First().position.Y > 800 {
			f.Entity[i].Pop()
		}
	}
	last := []*Entity{}
	for i := range f.Entity {
		last = append(last, f.Entity[i].First())
	}
	return last
}

func (f *Factory) Pop(index int) {
	f.Entity[index].Pop()
}

func (f *Factory) Draw(scene *ebiten.Image) {
	for i := range f.Entity {
		f.Entity[i].For(func(item *Entity) {
			item.Draw(scene)
		})
	}
}
