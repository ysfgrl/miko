package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ysfgrl/miko/miko/core"
	"github.com/ysfgrl/miko/miko/input"
	"path"
)

type EntityStatus string

const (
	Loading   EntityStatus = "loading"
	Idle                   = "idle"
	WalkLeft               = "walkLeft"
	WalkRight              = "walkRight"
	WalkDown               = "walkDown"
	WalkUp                 = "walkUp"
	RunLeft                = "runLeft"
	RunRight               = "runRight"
	RunDown                = "runDown"
	RunUp                  = "runUp"
)

var statusList = []EntityStatus{
	Idle,
	WalkLeft,
	WalkRight,
	WalkUp,
	WalkDown,
	RunUp,
	RunDown,
	RunLeft,
	RunRight,
}

type Entity struct {
	position *core.Position
	speed    *core.Position
	assets   map[EntityStatus][]*ebiten.Image
	asset    string
	index    int
	timer    int
	status   EntityStatus
	action   map[input.Action][]Action
	init     bool
	Life     *Life
}

func NewEntity(asset string, start *core.Position) *Entity {
	var life *Life

	if asset == "images/medo/5" || asset == "images/medo/6" {
		life = NewLife(-1)
	} else {
		life = NewLife(1)
	}
	return &Entity{
		position: start,
		speed: &core.Position{
			X: 0,
			Y: 0,
		},
		assets: make(map[EntityStatus][]*ebiten.Image),
		status: Loading,
		Life:   life,
		init:   false,
		asset:  asset,
		action: make(map[input.Action][]Action),
	}
}

func (e *Entity) WalkLeft() {
	e.position.X -= core.SPEED
	if _, ok := e.assets[WalkLeft]; ok {
		e.status = WalkLeft
	}
}

func (e *Entity) WalkRight() {
	e.position.X += core.SPEED
	if _, ok := e.assets[WalkRight]; ok {
		e.status = WalkRight
	}
}

func (e *Entity) WalkDown() {
	e.position.Y += core.SPEED
	if _, ok := e.assets[WalkDown]; ok {
		e.status = WalkDown
	}
}

func (e *Entity) WalkUp() {
	e.position.X -= core.SPEED
	if _, ok := e.assets[WalkUp]; ok {
		e.status = WalkUp
	}
}

func (e *Entity) AddAction(action input.Action, actions ...Action) {
	e.action[action] = actions
}

func (e *Entity) load() {
	e.assets["loading"] = core.Loading
	e.status = Loading
	for _, key := range statusList {
		list, err := core.LoadAssets(path.Join(e.asset, string(key)))
		if err == nil {
			e.assets[key] = list
		}
	}
	if _, ok := e.assets[Idle]; !ok {
		e.status = Idle
	}
	e.init = true
}

func (e *Entity) Update() {
	if !e.init {
		e.load()
	}
	e.status = Idle
	e.timer++
	for key, actions := range e.action {
		if input.Current().IsAction(key) {
			for _, action := range actions {
				action(e)
			}
		}
	}
	e.index = (e.timer / 10) % len(e.assets[e.status])
}

func (e *Entity) GetPosition() *core.Position {
	return e.position
}

func (e *Entity) GetBox() *core.Box {
	w, h := e.assets[e.status][0].Size()
	return &core.Box{
		Left:   int(e.position.X),
		Top:    int(e.position.Y),
		Right:  int(e.position.X) + w,
		Bottom: int(e.position.Y) + h,
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.position.X, e.position.Y)
	if list, ok := e.assets[e.status]; ok && len(list) > 0 {
		screen.DrawImage(e.assets[e.status][e.index], op)
	}
}
