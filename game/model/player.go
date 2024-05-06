package model

import (
	"mvvasilev/last_light/util"

	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
)

type Player struct {
	id       uuid.UUID
	position util.Position

	inventory *Inventory
}

func CreatePlayer(x, y int) *Player {
	p := new(Player)

	p.id = uuid.New()
	p.position = util.PositionAt(x, y)
	p.inventory = CreateInventory(util.SizeOf(8, 4))

	return p
}

func (p *Player) UniqueId() uuid.UUID {
	return p.id
}

func (p *Player) Position() util.Position {
	return p.position
}

func (p *Player) MoveTo(newPos util.Position) {
	p.position = newPos
}

func (p *Player) Presentation() (rune, tcell.Style) {
	return '@', tcell.StyleDefault
}

func (p *Player) Passable() bool {
	return false
}

func (p *Player) Transparent() bool {
	return false
}

func (p *Player) Inventory() *Inventory {
	return p.inventory
}

func (p *Player) Input(e *tcell.EventKey) {
}

func (p *Player) Tick(dt int64) {

}
